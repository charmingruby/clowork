package client

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type Client struct {
	apiClient    pb.ChatAPIClient
	streamClient pb.ChatStreamClient
}

func New(conn *grpc.ClientConn) *Client {
	apiClient := pb.NewChatAPIClient(conn)
	streamClient := pb.NewChatStreamClient(conn)

	return &Client{
		apiClient:    apiClient,
		streamClient: streamClient,
	}
}
