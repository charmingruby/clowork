package stream

import (
	"time"

	"github.com/charmingruby/clowork/api/proto/pb"
)

func (s *Server) handleHeartbeat(evt *pb.ClientEvent_Heartbeat) {
	payload := evt.Heartbeat

	if room, ok := s.rooms[payload.GetRoomId()]; ok {
		if sess, ok := room[payload.GetMemberId()]; ok {
			sess.lastBeat = time.Now()
		}
	}
}
