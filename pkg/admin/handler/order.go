package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/adminpb"
)

func ViewAllOrderHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	response, err := client.OrderHistory(ctx, &pb.AdminItemNoParams{})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
	}
	c.JSON(202, gin.H{
		"Status":  202,
		"Message": "Orderlist fetche successfully",
		"Data":    response,
	})
}

func ViewOrderHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	orderIDStr := c.Param("id")
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting orderID to int",
			"Error":   err.Error()})
		return
	}
	response, err := client.FindOrder(ctx, &pb.AdminItemID{ID: uint32(orderID)})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "Material fetched successfully",
		"Data":    response,
	})
}

func UserOrderHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in converting id to int",
			"Error":   err.Error()})
		return
	}
	response, err := client.FindOrdersByUser(ctx, &pb.AdminItemID{ID: uint32(userID)})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(202, gin.H{
		"Status":  202,
		"Message": "orders fetched successfully",
		"Data":    response,
	})
}
