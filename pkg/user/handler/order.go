package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

func PlaceOrderHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var order dto.Order

	if err := c.BindJSON(&order); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	id, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while user id converting",
			"Error":   ""})
		return
	}
	fmt.Println("userid", userID)

	response, err := client.PlaceOrder(ctx, &pb.UserOrder{
		User_ID:  uint32(userID),
		Item_ID:  uint32(order.ItemID),
		Quantity: int32(order.Quantity),
	})

	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"Status":  200,
		"Message": "Order placed successfully",
		"Data":    response,
	})

}

func ViewOrderHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	orderIDString := c.Param("id")
	orderID, err := strconv.Atoi(orderIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting orderID to int",
			"Error":   err.Error()})
		return
	}

	response, err := client.FindOrder(ctx, &pb.UserItemID{ID: uint32(orderID)})
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

func ViewAllOrderHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	// orderIDString := c.Param("id")
	// orderID, err := strconv.Atoi(orderIDString)
	// if err != nil {
	// 	c.AbortWithStatusJSON(400, gin.H{"Status": 400,
	// 		"Message": "error while converting categroyID to int",
	// 		"Error":   err.Error()})
	// 	return
	// }

	response, err := client.OrderHistory(ctx, &pb.NoParam{})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "OrderList fetched successfully",
		"Data":    response,
	})
}

func ViewAllOrderByUserHandler(c *gin.Context, client pb.UserServiceClient) {
	//set time out for the gRPC call
	timeOut := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	//get the userID from the context
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error while userid from context",
			"Error":   ""})
		return
	}
	userIDUint, ok := userID.(uint)
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error while user id converting",
			"Error":   "",
		})
		return
	}

	resp, err := client.FindOrdersByUser(ctx, &pb.UserItemID{ID: uint32(userIDUint)})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch items: " + err.Error()})
		return
	}

	if len(resp.Orders) == 0 {
		c.JSON(200, gin.H{"message": "No Orders found for this user", "ordres": []string{}})
		return
	}

	type Order struct {
		UserID    uint   `json:"user_id"`
		ItemID    uint   `json:"item_id"`    // The item being ordered
		Quantity  int    `json:"quantity"`   // Order quantity
		Status    string `json:"status"`     // Order status (Pending, Completed, etc.)
		CustomCut string `json:"custom_cut"` // Custom cut or any other preferences
		//IsCustom  bool    `json:"iscustom"`   // If the item has a custom size
		Amount    float64 `json:"amount"`
		PaymentID string  `json:"payment_id"`
	}

	orders := make([]Order, len(resp.Orders))
	for i, ord := range resp.Orders {
		orders[i] = Order{
			UserID:    uint(ord.Item_ID),
			ItemID:    uint(ord.Item_ID),
			Quantity:  int(ord.Quantity),
			Status:    ord.Status,
			Amount:    ord.Amount,
			PaymentID: ord.Payment_ID,
		}
	}

	c.JSON(200, gin.H{
		"Status": "List of orders",
		"Orders": orders})
}
