package handler

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

func UserPaymentHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	userIDstring := c.Query("id")
	orderIDString := c.Query("order_id")

	userID, err := strconv.Atoi(userIDstring)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in converting userID to int",
			"Error":   err.Error(),
		})
		return
	}

	orderID, err := strconv.Atoi(orderIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in converting orderID to int",
			"Error":   err.Error(),
		})
		return
	}

	response, err := client.UserCreatePayment(ctx, &pb.UserOrder{
		User_ID:  uint32(userID),
		Order_ID: uint32(orderID),
	})

	// Check for any errors in the client response
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}

	// Check if response is nil
	if response == nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Status":  500,
			"Message": "Payment response is nil",
		})
		return
	}

	// Log the response (for debugging purposes)
	log.Println("User create response:", response)

	// Proceed with rendering the HTML page
	c.HTML(200, "stripe.html", gin.H{
		"userID":    userID,
		"orderID":   orderID,
		"paymentID": response.PaymentId,
		"amount":    response.Amount,
		"client":    response.ClientSecret,
	})
}

// Define a struct for the expected request body
type PaymentRequest struct {
	UserID       string  `json:"user_id"`
	OrderID      string  `json:"order_id"`
	PaymentID    string  `json:"paymentID"`
	ClientSecret string  `json:"clientSecret"`
	Amount       float64 `json:"amount"`
}

func UserPaymentSuccessHandler(c *gin.Context, client pb.UserServiceClient) {
	// Set a timeout for the request context
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	// Bind the JSON data from the request body to the struct
	var req PaymentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Invalid request body",
			"Error":   err.Error(),
		})
		return
	}
	log.Println("response data", req)

	// Convert userID and orderID from string to int
	userID, err := strconv.Atoi(req.UserID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Error in converting userID to int",
			"Error":   err.Error(),
		})
		return
	}

	orderID, err := strconv.Atoi(req.OrderID)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Error in converting orderID to int",
			"Error":   err.Error(),
		})
		return
	}

	log.Println(orderID)
	// Call the client method to process the payment
	_, err = client.UserPaymentSuccess(ctx, &pb.UserPayment{
		User_ID:    uint32(userID),
		Payment_ID: req.PaymentID,
		Amount:     req.Amount,
		Order_ID:   uint32(orderID),
	})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "Error in client response",
			"Error":   err.Error(),
		})
		return
	}

	// Send a successful response
	c.JSON(200, gin.H{
		"status": true,
	})
}

func PaymentSuccessPage(ctx *gin.Context, client pb.UserServiceClient) {
	// Extract the "payment" query parameter
	paymentID := ctx.Query("paymentID")

	// Render the success page with the payment ID
	ctx.HTML(200, "success.html", gin.H{
		"paymentID": paymentID,
	})
}
