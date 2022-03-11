package grpcclient

import (
	"log"

	"github.com/robowealth-mutual-fund/blueprint-roa-golang/internals/config"
	api_v1 "github.com/robowealth-mutual-fund/blueprint-roa-golang/pkg/api/v1"

	"google.golang.org/grpc"
)

// HTTPGRPCClient ...
type HTTPGRPCClient struct {
	Config   config.Configuration
	PingPong api_v1.PingPongServiceClient
}

// Connect ...
func (client *HTTPGRPCClient) Connect() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("0.0.0.0:3000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("Connect to service on", "0.0.0.0:3000")

	client.PingPong = api_v1.NewPingPongServiceClient(conn)
}

// NewHTTPGRPCClient ...
func NewHTTPGRPCClient(config config.Configuration) *HTTPGRPCClient {
	// fmt.Printf("new grpc client: %+v\n", config)
	grpcClient := HTTPGRPCClient{
		Config: config,
	}

	grpcClient.Connect()

	return &grpcClient
}
