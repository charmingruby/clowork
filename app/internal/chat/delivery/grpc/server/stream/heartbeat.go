package stream

import (
	"time"

	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) heartbeatLoop() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for roomID := range s.rooms {
			s.broadcast(
				&pb.ServerEvent{
					EventSeq: 0,
					Event: &pb.ServerEvent_Heartbeat{
						Heartbeat: &pb.Heartbeat{
							RoomId: roomID,
							BeatAt: timestamppb.Now(),
						},
					},
				},
				roomID,
				"",
			)
		}
	}
}

func (s *Server) monitorHeartbeats() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()

		for roomID, members := range s.rooms {
			for memberID, sess := range members {
				if now.Sub(sess.lastBeat) > 20*time.Second {
					delete(members, memberID)

					s.broadcast(
						&pb.ServerEvent{
							EventSeq: 0,
							Event: &pb.ServerEvent_RoomLeft{
								RoomLeft: &pb.RoomLeft{
									RoomId:   roomID,
									Nickname: sess.nickname,
								},
							},
						},
						roomID,
						memberID,
					)
				}
			}
		}
	}
}
