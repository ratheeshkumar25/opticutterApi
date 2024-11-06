package handler

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/adminpb"
)

func ViewAllItemHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	resp, err := client.FindAllItem(ctx, &pb.AdminItemNoParams{})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(202, gin.H{
		"Status":  202,
		"Message": "Items fetched successfully",
		"Data":    resp,
	})
}
