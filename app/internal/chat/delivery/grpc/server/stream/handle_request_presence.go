package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

func (s *Server) handleRequestPresence(
	stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent],
	evt *pb.ClientEvent_RequestPresence,
) error {
	payload := evt.RequestPresence

	var presences []*pb.Presence
	if room := s.rooms[payload.GetRoomId()]; room != nil {
		for _, sess := range room {
			presences = append(presences, &pb.Presence{
				Nickname: sess.nickname,
				Hostname: sess.hostname,
			})
		}
	}

	s.sendTo(
		&pb.ServerEvent{
			EventSeq: 0,
			Event: &pb.ServerEvent_RoomPresence{
				RoomPresence: &pb.RoomPresence{
					Presences: presences,
				},
			},
		},
		nil,
		payload.GetRoomId(),
		payload.GetSenderId(),
	)

	return stream.Send(&pb.ServerEvent{
		EventSeq: 0,
		Event: &pb.ServerEvent_RoomPresence{
			RoomPresence: &pb.RoomPresence{
				Presences: presences,
			},
		},
	})
}
