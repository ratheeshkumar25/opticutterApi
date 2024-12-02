package chat

import (
	"fmt"
	"log"

	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/pb"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial function setup the connection between chat service and api gateway
func ClientDial(cfg config.Config) (pb.ChatServiceClient, error) {
	grpcAddr := fmt.Sprintf("localhost:%s", cfg.CHATPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}

	log.Printf("Succesfully connected to chat client at port: %v", cfg.CHATPORT)
	return pb.NewChatServiceClient(grpc), nil
}
