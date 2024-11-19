package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

func GetCuttingResultHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	itemIDString := c.Param("id")
	itemID, err := strconv.Atoi(itemIDString)
	if err != nil {
		c.Copy().AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error while converting itemID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.UserGetCuttingResult(ctx, &pb.UserItemID{ID: uint32(itemID)})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "cutting result fetched successfully",
		"Data":    response,
	})
}
