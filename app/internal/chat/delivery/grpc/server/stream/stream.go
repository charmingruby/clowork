package stream

import (
	"context"
	"errors"
	"io"
	"sync/atomic"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

func (s *Server) Stream(stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]) error {
	ctx := context.Background()

	for {
		cevt, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return err
		}

		switch evt := cevt.GetEvent().(type) {
		case *pb.ClientEvent_JoinRoom:
			err := s.handleJoinRoom(ctx, stream, evt)
			if err != nil {
				s.log.Error("handle join room error", "error", err.Error())
			}

			if err := s.sendAck(stream, cevt.GetClientMsgId(), err); err != nil {
				s.log.Error("ack error", "error", err.Error())
			}
		case *pb.ClientEvent_LeaveRoom:
			err := s.handleLeaveRoom(ctx, evt)
			if err != nil {
				s.log.Error("handle leave room error", "error", err.Error())
			}

			if err := s.sendAck(stream, cevt.GetClientMsgId(), err); err != nil {
				s.log.Error("ack error", "error", err.Error())
			}
		case *pb.ClientEvent_SendMessage:
			err := s.handleSendMessage(ctx, evt)
			if err != nil {
				s.log.Error("handle send message error", "error", err.Error())
			}

			if err := s.sendAck(stream, cevt.GetClientMsgId(), err); err != nil {
				s.log.Error("ack error", "error", err.Error())
			}
		case *pb.ClientEvent_RequestPresence:
			s.handleRequestPresence(evt)

			if err := s.sendAck(stream, cevt.GetClientMsgId(), nil); err != nil {
				s.log.Error("ack error", "error", err.Error())
			}
		case *pb.ClientEvent_Heartbeat:
			s.handleHeartbeat(evt)
		}
	}
}

func (s *Server) sendAck(
	stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent],
	clientMsgID string,
	err error,
) error {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return stream.Send(&pb.ServerEvent{
		EventSeq: s.nextSeq(),
		Event: &pb.ServerEvent_Ack{
			Ack: &pb.Ack{
				ClientMsgId: clientMsgID,
				Success:     err == nil,
				Error:       errMsg,
			},
		},
	})
}

func (s *Server) broadcast(evt *pb.ServerEvent, roomID, targetMemberID string) {
	hasTarget := targetMemberID != ""

	if room := s.rooms[roomID]; room != nil {
		for memberID, sess := range room {
			if !hasTarget || (hasTarget && memberID != targetMemberID) {
				if err := sess.stream.Send(evt); err != nil {
					s.log.Error("broadcast error",
						"error", err.Error(),
						"room_id", roomID,
						"member_id", memberID,
						"event", evt.GetEvent(),
					)
				}
			}
		}
	}
}

func (s *Server) sendTo(
	evt *pb.ServerEvent,
	customStream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent],
	roomID,
	memberID string,
) {
	if customStream != nil {
		if err := customStream.Send(evt); err != nil {
			s.log.Error("broadcast error",
				"error", err.Error(),
				"room_id", roomID,
				"member_id", memberID,
				"event", evt.GetEvent(),
			)
		}

		return
	}

	sess := s.rooms[roomID][memberID]
	if err := sess.stream.Send(evt); err != nil {
		s.log.Error("broadcast error",
			"error", err.Error(),
			"room_id", roomID,
			"member_id", memberID,
			"event", evt.GetEvent(),
		)
	}
}

func (s *Server) nextSeq() uint64 {
	return atomic.AddUint64(&s.eventSeq, 1)
}
