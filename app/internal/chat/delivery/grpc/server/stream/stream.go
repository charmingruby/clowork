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
		}
	}
}

func (s *Server) broadcast(evt *pb.ServerEvent, roomID, excludedMemberID string) {
	hasExclusion := excludedMemberID != ""

	if room := s.rooms[roomID]; room != nil {
		for memberID, sess := range room {
			if !hasExclusion || (hasExclusion && memberID != excludedMemberID) {
				if err := sess.stream.Send(evt); err != nil {
					s.log.Error("broadcast error",
						"error", err.Error(),
						"room_id", room,
						"member_id", memberID,
						"event", evt.GetEvent(),
					)
				}
			}
		}
	}
}
