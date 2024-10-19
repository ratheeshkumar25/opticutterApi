package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

// ViewProfileHandler handles the user profile view request.
func ViewProfileHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	response, err := client.ViewProfile(ctx, &pb.ID{
		ID: uint32(userID),
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "profile fetched successfully",
		"Data":    response,
	})
}

// EditProfileHandler handles the user profile edit request.
func EditProfileHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var user dto.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	response, err := client.EditProftle(ctx, &pb.Profile{
		User_ID:    uint32(userID),
		First_Name: user.FirstName,
		Last_Name:  user.LastName,
		Phone:      user.Phone,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Profile edited succesfully",
		"Data":    response,
	})
}

// ChangePasswordHandler handles the user password change request.
func ChangePasswordHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var user dto.Password

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}

	response, err := client.ChangePassword(ctx, &pb.Password{
		User_ID:          uint32(userID),
		Old_Password:     user.Old,
		New_Password:     user.New,
		Confirm_Password: user.Confirm,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "PAssword updated succesfully",
		"Data":    response,
	})
}
