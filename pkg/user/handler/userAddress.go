package handler

import (
	"context"
	"log"

	//"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

// AddAddressHandler handles the request for adding address of user
func AddAddressHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var address dto.Address

	if err := c.BindJSON(&address); err != nil {
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

	response, err := client.AddAddress(ctx, &pb.Address{
		User_ID: uint32(userID),
		House:   address.House,
		Street:  address.Street,
		City:    address.City,
		Zip:     uint32(address.ZIP),
		State:   address.State,
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
		"Message": "address added successfully",
		"Data":    response,
	})
}

// EditAddressHandler handles the request for editing address of user
func EditAddressHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	addressIDString := c.Param("id")
	addressID, err := strconv.Atoi(addressIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting addressID to int",
			"Error":   err.Error()})
		return
	}

	var address dto.Address
	if err := c.BindJSON(&address); err != nil {
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

	response, err := client.EditAddress(ctx, &pb.Address{
		ID:      uint32(addressID),
		User_ID: uint32(userID),
		House:   address.House,
		Street:  address.Street,
		City:    address.City,
		Zip:     uint32(address.ZIP),
		State:   address.State,
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
		"Message": "address edited successfully",
		"Data":    response,
	})
}

// ViewAllAddressHandler handles fetching the address of the authenticated user
func ViewAllAddressHandler(c *gin.Context, client pb.UserServiceClient) {
	// Set a timeout for the gRPC call
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	// Get the user ID from the context (set by the JWT middleware)
	userID, exists := c.Get("user_id")
	log.Print("userID", userID)
	if !exists {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while user id from context",
			"Error":   ""})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error while user id converting",
			"Error":   "",
		})
		return
	}

	// Call the gRPC service with the user ID
	resp, err := client.ViewAllAddress(ctx, &pb.ID{ID: uint32(userIDUint)}) // Convert to uint32
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch addresses: " + err.Error()})
		return
	}
	log.Printf("address Response: %+v", resp)

	// Check if the response contains addresses
	if len(resp.Addresses) == 0 {
		// Return a message indicating no addresses
		c.JSON(http.StatusOK, gin.H{"message": "No addresses found for this user", "addresses": []string{}})
		return
	}

	// Convert the gRPC response to a more API-friendly format
	type Address struct {
		ID     uint32 `json:"id"`
		House  string `json:"house"`
		Street string `json:"street"`
		City   string `json:"city"`
		Zip    uint32 `json:"zip"`
		State  string `json:"state"`
		UserID uint32 `json:"user_id"`
	}

	addresses := make([]Address, len(resp.Addresses))
	for i, addr := range resp.Addresses {
		addresses[i] = Address{
			ID:     addr.ID,
			House:  addr.House,
			Street: addr.Street,
			City:   addr.City,
			Zip:    addr.Zip,
			State:  addr.State,
			UserID: addr.User_ID,
		}
	}

	// Return the response as JSON with the list of addresses
	c.JSON(http.StatusOK, gin.H{"addresses": addresses})
}

// RemoveAddressHandler handles the request for removing address of user
func RemoveAddressHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	addressIDString := c.Param("id")
	addressID, err := strconv.Atoi(addressIDString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while converting addressID to int",
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

	response, err := client.RemoveAddress(ctx, &pb.IDs{
		ID:      uint32(addressID),
		User_ID: uint32(userID),
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
		"Message": "address edited successfully",
		"Data":    response,
	})
}
