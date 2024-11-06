package handler

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/adminpb"
)

func AddMaterialHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	var material dto.Material

	if err := c.BindJSON(&material); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "erroe while binding",
			"Error":   err.Error()})
		return
	}
	response, err := client.AddMaterial(ctx, &pb.AdminMaterial{
		Material_Name: material.Name,
		Description:   material.Description,
		Stock:         int32(material.Stock),
		Price:         material.Price,
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
		"Message": "material added successfully",
		"Data":    response})
}

func UpdateMaterialHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()
	materialString := c.Param("id")
	materialID, err := strconv.Atoi(materialString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while converting itemID to int",
			"Error":   err.Error()})
		return
	}
	var material dto.Material

	if err := c.BindJSON(&material); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.EditMaterial(ctx, &pb.AdminMaterial{
		Material_ID:   uint32(materialID),
		Material_Name: material.Name,
		Description:   material.Description,
		Stock:         int32(material.Stock),
		Price:         material.Price,
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
		"Message": "material details updated successfully",
		"Data":    response})

}

func ViewMaterialHandler(c *gin.Context, client pb.AdminServiceClient) {
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

	response, err := client.FindMaterialByID(ctx, &pb.AdminMaterialID{ID: uint32(materialID)})
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

func ViewAllMaterialHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	response, err := client.FindAllMaterial(ctx, &pb.AdminItemNoParams{})
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"Status": 400,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"Status":  201,
		"Message": "Materials fetched successfully",
		"Data":    response,
	})
}

func RemoveMaterialHandler(c *gin.Context, client pb.AdminServiceClient) {
	timeOut := time.Second * 100
	ctx, cancel := context.WithTimeout(c, timeOut)
	defer cancel()

	materialIDString := c.Param("id")
	materialID, err := strconv.Atoi(materialIDString)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"Status":  400,
			"Message": "error in while converting materialID to integer",
			"Error":   err.Error(),
		})
		return
	}

	//call the gRPC service to remove material
	response, err := client.RemoveMaterial(ctx, &pb.AdminMaterialID{ID: uint32(materialID)})
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"Status":  500,
			"Message": "error in client response",
			"Error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"Status":  200,
		"Message": "Material removed successfully",
		"Data":    response,
	})
}
