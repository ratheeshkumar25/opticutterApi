package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/handler"
)

// UserSignup handles the user signup request.
func (u *User) UserSignup(ctx *gin.Context) {
	handler.UserSignupHandler(ctx, u.client)
}

// UserVerify handles the user verification request.
func (u *User) UserVerify(ctx *gin.Context) {
	handler.VerificationHandler(ctx, u.client)
}

// UserLogin handles the user login request.
func (u *User) UserLogin(ctx *gin.Context) {
	handler.UserLoginHandler(ctx, u.client)
}

// AddAddress handles the request to add a new address for the user.
func (u *User) AddAddress(ctx *gin.Context) {
	handler.AddAddressHandler(ctx, u.client)
}

// EditAddress handles the request to edit an existing address for the user.
func (u *User) EditAddress(ctx *gin.Context) {
	handler.EditAddressHandler(ctx, u.client)
}

// ViewAddress handles the address fetch for user
func (u *User) ViewAllAddress(ctx *gin.Context) {
	handler.ViewAllAddressHandler(ctx, u.client)
}

// RemoveAddress handles the request to remove an existing address for the user.
func (u *User) RemoveAddress(ctx *gin.Context) {
	handler.RemoveAddressHandler(ctx, u.client)
}

// ViewProfile handles the request to view the user's profile.
func (u *User) ViewProfile(ctx *gin.Context) {
	handler.ViewProfileHandler(ctx, u.client)
}

// EditProfile handles the request to edit the user's profile.
func (u *User) EditProfile(ctx *gin.Context) {
	handler.EditProfileHandler(ctx, u.client)
}

// ChangePassword handles the request to change the user's password.
func (u *User) ChangePassword(ctx *gin.Context) {
	handler.ChangePasswordHandler(ctx, u.client)
}
