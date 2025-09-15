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
	payload := evt.JoinRoom

	if _, isRoomRegistered := s.rooms[payload.GetRoomId()]; !isRoomRegistered {
		s.rooms[payload.GetRoomId()] = map[string]*session{}
	}

	if s.rooms[payload.GetRoomId()] == nil {
		s.rooms[payload.GetRoomId()] = make(map[string]*session)
	}

	memberID, err := s.usecase.JoinRoom(ctx, usecase.JoinRoomInput{
		Nickname: payload.GetNickname(),
		Hostname: payload.GetHostname(),
		RoomID:   payload.GetRoomId(),
	})
	if err != nil {
		return err
	}

	sess := &session{
		memberID: memberID,
		nickname: payload.GetNickname(),
		stream:   stream,
	}

	s.rooms[payload.GetRoomId()][memberID] = sess

	if err := sess.stream.Send(&pb.ServerEvent{
		EventSeq: 0,
		Event: &pb.ServerEvent_RoomJoined{
			RoomJoined: &pb.RoomJoined{
				RoomId:   payload.GetRoomId(),
				MemberId: memberID,
				Nickname: payload.GetNickname(),
			},
		},
	}); err != nil {
		return err
	}

	s.broadcast(&pb.ServerEvent{
		EventSeq: 0,
		Event: &pb.ServerEvent_RoomJoined{
			RoomJoined: &pb.RoomJoined{
				RoomId:   payload.GetRoomId(),
				MemberId: memberID,
				Nickname: payload.GetNickname(),
			},
		},
	},
		payload.GetRoomId(),
		memberID,
	)

	return nil
}
