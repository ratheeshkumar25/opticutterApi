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

func (a *Admin) AddMaterial(ctx *gin.Context) {
	handler.AddMaterialHandler(ctx, a.Client)
}

func (a *Admin) FindMaterialByID(ctx *gin.Context) {
	handler.ViewMaterialHandler(ctx, a.Client)
}

func (a *Admin) FindAllMaterial(ctx *gin.Context) {
	handler.ViewAllMaterialHandler(ctx, a.Client)
}

func (a *Admin) EditMaterial(ctx *gin.Context) {
	handler.UpdateMaterialHandler(ctx, a.Client)
}

func (a *Admin) RemoveMaterial(ctx *gin.Context) {
	handler.RemoveMaterialHandler(ctx, a.Client)
}

func (a *Admin) FindAllItem(ctx *gin.Context) {
	handler.ViewAllItemHandler(ctx, a.Client)
}

func (a *Admin) OrderHistory(ctx *gin.Context) {
	handler.ViewAllOrderHandler(ctx, a.Client)
}

func (a *Admin) FindOrder(ctx *gin.Context) {
	handler.ViewOrderHandler(ctx, a.Client)
}

func (a *Admin) FindOrdersByUser(ctx *gin.Context) {
	handler.UserOrderHandler(ctx, a.Client)
}

// // AdminLogin handles the admin login request in net/http.
// func (a *Admin) AdminLogin(w http.ResponseWriter, r *http.Request) {
// 	handler.AdminLoginHandler(w, r, a.Client)
// }

// // BlockUser handles the admin request to block a user.
// func (a *Admin) BlockUser(w http.ResponseWriter, r *http.Request) {
// 	handler.BlockUserHandler(w, r, a.Client)
// }
