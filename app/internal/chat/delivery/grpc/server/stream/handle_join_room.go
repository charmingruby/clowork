package stream

import (
	"context"

	"github.com/charmingruby/clowork/api/proto/pb"
	"github.com/charmingruby/clowork/internal/chat/usecase"
	"google.golang.org/grpc"
)

func (s *Server) handleJoinRoom(
	ctx context.Context,
	stream grpc.BidiStreamingServer[pb.ClientEvent, pb.ServerEvent],
	evt *pb.ClientEvent_JoinRoom,
) error {
	joinRoom := evt.JoinRoom

	if _, isRoomRegistered := s.rooms[joinRoom.RoomId]; !isRoomRegistered {
		s.rooms[joinRoom.RoomId] = map[string]*session{}
	}

	if s.rooms[joinRoom.RoomId] == nil {
		s.rooms[joinRoom.RoomId] = make(map[string]*session)
	}

	memberID, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
		Nickname: joinRoom.Nickname,
		Hostname: joinRoom.Hostname,
		RoomID:   joinRoom.RoomId,
	})
	if err != nil {
		return err
	}

	sess := &session{
		memberID: memberID,
		stream:   stream,
	}

	s.rooms[joinRoom.RoomId][memberID] = sess

	s.broadcastToRoom(&pb.ServerEvent{
		EventSeq: 0,
		Event: &pb.ServerEvent_RoomJoined{
			RoomJoined: &pb.RoomJoined{
				RoomId:   joinRoom.RoomId,
				MemberId: memberID,
				Nickname: joinRoom.Nickname,
			},
		},
	},
		joinRoom.RoomId,
		memberID,
	)

	return nil
}
