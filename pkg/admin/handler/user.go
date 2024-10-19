package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/adminpb"
)

// BlockUserHandler function will send block user request to client.
func BlockUserHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	UserIDString := c.Param("id")
	UserID, err := strconv.Atoi(UserIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting userID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.AdminBlockUser(ctx, &pb.AdID{
		ID: uint32(UserID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "User blocked successfully",
		"Data":    response,
	})
}

// BlockUserHandler function will send block user request to client.
func UnblockUserHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	UserIDString := c.Param("id")
	UserID, err := strconv.Atoi(UserIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting userID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.AdminUnblockUser(ctx, &pb.AdID{
		ID: uint32(UserID),
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "User unblocked successfully",
		"Data":    response,
	})
}
