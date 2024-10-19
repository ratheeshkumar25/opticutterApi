package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/handler"
)

// AdminLogin handles the admin login request.
func (a *Admin) AdminLogin(ctx *gin.Context) {
	handler.AdminLoginHandler(ctx, a.Client)
}

func (a *Admin) BlockUser(ctx *gin.Context) {
	handler.BlockUserHandler(ctx, a.Client)
}

func (a *Admin) UnblockUser(ctx *gin.Context) {
	handler.UnblockUserHandler(ctx, a.Client)
}

// // AdminLogin handles the admin login request in net/http.
// func (a *Admin) AdminLogin(w http.ResponseWriter, r *http.Request) {
// 	handler.AdminLoginHandler(w, r, a.Client)
// }

// // BlockUser handles the admin request to block a user.
// func (a *Admin) BlockUser(w http.ResponseWriter, r *http.Request) {
// 	handler.BlockUserHandler(w, r, a.Client)
// }
