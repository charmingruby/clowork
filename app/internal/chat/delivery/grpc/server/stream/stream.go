package stream

import (
	"context"
	"errors"
	"io"

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
			if err := s.handleJoinRoom(ctx, stream, evt); err != nil {
				s.log.Error("handle join room error", "error", err.Error())
				continue
			}
		case *pb.ClientEvent_LeaveRoom:
			if err := s.handleLeaveRoom(ctx, evt); err != nil {
				s.log.Error("handle leave room error", "error", err.Error())
				continue
			}
		case *pb.ClientEvent_SendMessage:
			if err := s.handleSendMessage(ctx, evt); err != nil {
				s.log.Error("handle send message error", "error", err.Error())
				continue
			}
		case *pb.ClientEvent_RequestPresence:
			if err := s.handleRequestPresence(stream, evt); err != nil {
				s.log.Error("handle request presence error", "error", err.Error())
				continue
			}
		}
	}
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
