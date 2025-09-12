package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type session struct {
	nickname      string
	hostname      string
	currentRoomID string
}

type Client struct {
	streamClient pb.ChatStreamClient

	stream  grpc.BidiStreamingClient[pb.ClientEvent, pb.ServerEvent]
	session *session
	console chan string
}

func New(conn *grpc.ClientConn, console chan string) *Client {
	streamClient := pb.NewChatStreamClient(conn)

	return &Client{
		streamClient: streamClient,
		stream:       nil,
		console:      console,
		session: &session{
			nickname:      "",
			hostname:      "",
			currentRoomID: "",
		},
	}
}
