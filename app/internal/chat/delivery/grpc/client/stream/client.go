package stream

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type session struct {
	nickname      string
	hostname      string
	currentRoomID string
	memberID      string
}

type Client struct {
	streamClient pb.ChatStreamClient

	stream       grpc.BidiStreamingClient[pb.ClientEvent, pb.ServerEvent]
	session      *session
	msgCh        chan string
	joinedCh     chan struct{}
	lastEventSeq uint64
}

func New(conn *grpc.ClientConn, msgCh chan string) *Client {
	streamClient := pb.NewChatStreamClient(conn)

	return &Client{
		streamClient: streamClient,
		stream:       nil,
		msgCh:        msgCh,
		session: &session{
			nickname:      "",
			hostname:      "",
			currentRoomID: "",
			memberID:      "",
		},
	}
}
