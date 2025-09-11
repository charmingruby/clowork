package server

import (
	"io"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

func (s *Server) Stream(stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent]) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.log.Info("event received",
			"client_msg_id", in.GetClientMsgId(),
			"last_event_seq", in.GetLastEventSeq(),
			"event", in.GetEvent(),
		)

		if err := stream.Send(&pb.ServerEvent{
			EventSeq: 1,
			Event: &pb.ServerEvent_Ack{
				Ack: &pb.Ack{Id: in.ClientMsgId},
			},
		}); err != nil {
			return err
		}
	}
}
