package di

import (
	"log"

	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/server"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user"
)

func Init() {
	cnf, err := config.LoadConfig()
	if err != nil {
		log.Printf("Error loading configuration file")
	}

	// Initialize the server
	srv := server.NewServer()
	//srv.R.LoadHTMLGlob("template/*")
	srv.R.LoadHTMLGlob("/home/user/Desktop/opti_cut_Api_gateway/template/*")
	//srv.R.LoadHTMLFiles("template/stripe.html")

	// Use the HTTPServer field instead of calling it
	user.NewUserRoute(srv.R, *cnf)
	admin.NewAdminRoute(srv.R, *cnf)
	srv.StartServer(cnf.APIPORT)
}
