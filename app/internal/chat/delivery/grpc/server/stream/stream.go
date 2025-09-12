package stream

import (
	"errors"
	"io"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/pkg/core"
	"google.golang.org/grpc"
)

func (s *Server) Stream(stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]) error {
	ctx := stream.Context()

	s.stream = stream

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
				s.handleErr(err)
				continue
			}
		}
	}
}

func (s *Server) handleErr(err error) {
	var databaseErr *core.DatabaseError
	if errors.As(err, &databaseErr) {
		s.log.Error("stream error", "reason", "database", "error", databaseErr.Unwrap().Error())
		return
	}

	s.log.Error("stream error", "reason", "unknown", "error", err.Error())
}
