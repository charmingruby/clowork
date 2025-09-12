package stream

import (
	"io"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

func (s *Server) Stream(stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]) error {
	ctx := stream.Context()

	for {
		evt, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.log.Info("event received",
			"client_msg_id", evt.GetId(),
			"event", evt.GetEvent(),
		)

		switch e := evt.Event.(type) {
		case *pb.ClientEvent_JoinRoom:
			if err := s.handleJoinRoom(ctx, e); err != nil {
				s.log.Error("JoinRoom stream event processing error", "error", err.Error())
				continue
			}
		}
	}
}
