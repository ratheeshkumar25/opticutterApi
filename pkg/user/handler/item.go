package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

func AddItemHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	var item dto.Item
	if err := c.BindJSON(&item); err != nil {
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

	response, err := client.AddItem(ctx, &pb.UserItem{
		User_ID:       uint32(userID),
		Item_Name:     item.ItemName,
		Material_ID:   uint32(item.MaterialID),
		Length:        uint32(item.Length),
		Width:         uint32(item.Width),
		Fixed_Size_ID: uint32(item.FixedSizeID),
		Is_Custom:     item.IsCustom,
	})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "item added successfully",
		"Data":    response})

}

func EditItemHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	itemIDString := c.Param("id")
	itemID, err := strconv.Atoi(itemIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting itemID to int",
			"Error":   err.Error()})
		return
	}

	var item dto.Item

	if err := c.BindJSON(&item); err != nil {
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

	response, err := client.EditItem(ctx, &pb.UserItem{
		User_ID:       uint32(userID),
		Item_ID:       uint32(itemID),
		Item_Name:     item.ItemName,
		Material_ID:   uint32(item.MaterialID),
		Length:        uint32(item.Length),
		Width:         uint32(item.Width),
		Fixed_Size_ID: uint32(item.FixedSizeID),
		Is_Custom:     item.IsCustom,
	})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Data":    response,
			"Error":   err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "item added successfully",
		"Data":    response})
}

func ViewAllItemHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	userID, exists := c.Get("user_id")
	log.Print("userID", userID)
	if !exists {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while user id from context",
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

	fmt.Println("userid", userIDUint)

	// Call the gRPC service with the user ID
	resp, err := client.FindAllItem(ctx, &pb.NoParam{}) // Convert to uint32
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}
	log.Printf("address Response: %+v", resp)

	// Check if the response contains addresses
	if len(resp.Items) == 0 {
		// Return a message indicating no addresses
		c.JSON(200, gin.H{"message": "No addresses found for this user", "addresses": []string{}})
		return
	}

	type Item struct {
		ItemName    string `json:"item_name"`
		MaterialID  uint   `json:"material_id"`
		Length      uint   `json:"length"`
		Width       uint   `json:"width"`
		FixedSizeID uint   `json:"fixed_size_id"`
		IsCustom    bool   `json:"default:false"` // Flag indicating if size is custom
		UserID      uint   `json:"user_id"`
	}

	items := make([]Item, len(resp.Items))
	for i, itm := range resp.Items {
		items[i] = Item{
			ItemName:    itm.Item_Name,
			MaterialID:  uint(itm.Material_ID),
			Length:      uint(itm.Length),
			Width:       uint(itm.Width),
			FixedSizeID: uint(itm.Fixed_Size_ID),
			IsCustom:    itm.Is_Custom,
			UserID:      uint(itm.User_ID),
		}
	}
	c.JSON(202, gin.H{
		"Status":  202,
		"Message": "Items fetched successfully",
		"Data":    resp,
	})

}

func ViewItemHandler(c *gin.Context, client pb.UserServiceClient) {

}

func RemoveItemHandler(c *gin.Context, client pb.UserServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	// Get item ID from URL params
	itemIDString := c.Param("id")
	itemID, err := strconv.Atoi(itemIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting itemID to int",
			"Error":   err.Error()})
		return
	}

	// Get the user_id from context
	userID, ok := c.Get("user_id")
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while fetching user id from context",
			"Error":   ""})
		return
	}

	userIDUint, ok := userID.(uint)
	if !ok {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting user id",
			"Error":   ""})
		return
	}
	fmt.Println("user", userIDUint)

	// Call the gRPC service to remove the item
	response, err := client.RemoveItem(ctx, &pb.UserItemID{
		ID: uint32(itemID),
		// User_ID: uint32(userIDUint),
	})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"Status": 500,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"Status":  200,
		"Message": "Item removed successfully",
		"Data":    response,
	})
}

func ViewAllItemByUserHandler(c *gin.Context, client pb.UserServiceClient) {
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

	resp, err := client.FindAllItemByUser(ctx, &pb.UserItemID{ID: uint32(userIDUint)})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch items: " + err.Error()})
		return
	}

	if len(resp.Items) == 0 {
		c.JSON(200, gin.H{"message": "No items found for this user", "items": []string{}})
		return
	}

	type Item struct {
		ItemName    string `json:"item_name"`
		ItemID      uint   `json:"item_id"`
		MaterialID  uint   `json:"material_id"`
		Length      uint   `json:"length"`
		Width       uint   `json:"width"`
		FixedSizeID uint   `json:"fixed_size_id"`
		//IsCustom    bool    `json:"default:false"`
		UserID   uint    `json:"user_id"`
		EstPrice float32 `json:"est_prcice"`
	}

	items := make([]Item, len(resp.Items))
	for i, itm := range resp.Items {
		items[i] = Item{
			ItemName:    itm.Item_Name,
			ItemID:      uint(itm.Item_ID),
			MaterialID:  uint(itm.Material_ID),
			Length:      uint(itm.Length),
			Width:       uint(itm.Width),
			FixedSizeID: uint(itm.Fixed_Size_ID),
			//IsCustom:    itm.Is_Custom,
			UserID:   uint(itm.User_ID),
			EstPrice: itm.Estimated_Price,
		}
	}
	c.JSON(200, gin.H{
		"Status": "Items orderby User",
		"Items":  items})
}
