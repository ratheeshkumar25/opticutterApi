package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

func MaterialByIDHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	materialIDString := c.Param("id")
	materialID, err := strconv.Atoi(materialIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting categroyID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.FindMaterialByID(ctx, &pb.UserMaterialID{ID: uint32(materialID)})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "Products fetched successfully",
		"Data":    response,
	})
}

func FindAllMaterialHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	response, err := client.FindAllMaterial(ctx, &pb.NoParam{})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "Products fetched successfully",
		"Data":    response,
	})
}
