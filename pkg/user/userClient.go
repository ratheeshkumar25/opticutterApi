package user

import (
	"fmt"
	"log"

	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial method connects to the service to the user client
func ClientDial(cfg config.Config) (pb.UserServiceClient, error) {
	grpcAddr := fmt.Sprintf("localhost:%s", cfg.USERPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to user client at address : %s", grpcAddr)
	return pb.NewUserServiceClient(grpc), nil

}
