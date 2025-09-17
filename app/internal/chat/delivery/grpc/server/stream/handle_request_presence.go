package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
)

func (s *Server) handleRequestPresence(evt *pb.ClientEvent_RequestPresence) {
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
			EventSeq: s.nextSeq(),
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
}
