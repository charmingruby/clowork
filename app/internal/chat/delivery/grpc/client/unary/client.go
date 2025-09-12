package unary

import (
	"github.com/charmingruby/clowork/api/proto/pb"
	"google.golang.org/grpc"
)

type Client struct {
	apiClient pb.ChatAPIClient
}

func New(conn *grpc.ClientConn) *Client {
	apiClient := pb.NewChatAPIClient(conn)

	return &Client{
		apiClient: apiClient,
	}
}
