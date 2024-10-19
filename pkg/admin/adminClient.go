package admin

import (
	"fmt"
	"log"

	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/adminpb"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ClientDial method connects to the service to the admin client
func ClientDial(cfg config.Config) (pb.AdminServiceClient, error) {
	grpcAddr := fmt.Sprintf("localhost:%s", cfg.ADMINPORT)
	grpc, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error dialing to grpc to client : %s", err.Error())
		return nil, err
	}
	log.Printf("Successfully connected to admin client at port : %s", cfg.USERPORT)
	return pb.NewAdminServiceClient(grpc), nil
}
