package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type session struct {
	nickname      string
	hostname      string
	currentRoomID string
	lastEventSeq  int
}

type Client struct {
	streamClient pb.ChatStreamClient

	stream  grpc.BidiStreamingClient[pb.ClientEvent, pb.ServerEvent]
	session *session
}

func New(conn *grpc.ClientConn) *Client {
	streamClient := pb.NewChatStreamClient(conn)

	return &Client{
		streamClient: streamClient,
		stream:       nil,
		session: &session{
			nickname:      "",
			hostname:      "",
			currentRoomID: "",
			lastEventSeq:  0,
		},
	}
}
