package stream

import (
	"io"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

func (s *Server) Stream(stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]) error {
	ctx := stream.Context()

	for {
		cevt, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		switch evt := cevt.Event.(type) {
		case *pb.ClientEvent_JoinRoom:
			if err := s.handleJoinRoom(ctx, stream, evt); err != nil {
				s.log.Error("handle join room error", "error", err.Error())
				continue
			}
		}
	}
}

func (s *Server) broadcastToRoom(evt *pb.ServerEvent, roomID, excludedMemberID string) {
	if room := s.rooms[roomID]; room != nil {
		for memberID, sess := range room {
			if memberID != excludedMemberID {
				if err := sess.stream.Send(evt); err != nil {
					s.log.Error("broadcast error",
						"room_id", room,
						"member_id", memberID,
						"event", evt.GetEvent(),
					)
				}
			}
		}
	}
}
