package grpc_clients

import (
	pb "gcjade/shared/proto/catalogue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CatalogueServiceClient struct {
	Client pb.CatalogueServiceClient
	conn   *grpc.ClientConn
}

func NewCatalogueServiceClient() (*CatalogueServiceClient, error) {
	serviceURL := "catalogue-service:9092"

	conn, err := grpc.NewClient(serviceURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewCatalogueServiceClient(conn)

	return &CatalogueServiceClient{
		Client: client,
		conn:   conn,
	}, nil
}

func (c *CatalogueServiceClient) Close() {
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			return
		}
	}
}
