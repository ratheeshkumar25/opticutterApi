package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/utils"
)

// UserSignupHandler handles the user signup request.
func UserSignupHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()
	var user dto.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	validate := validator.New()
	validate.RegisterValidation("email", utils.EmailValidation)
	validate.RegisterValidation("phone", utils.PhoneNumberValidation)
	err := validate.Struct(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "validation error",
			"Error":   err.Error()})

		for _, e := range err.(validator.ValidationErrors) {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
				"Message": fmt.Sprintf("Error in field %v, error: %v", e.Field(), e.Tag()),
				"Error":   e})
		}
	}

	response, err := client.UserSignup(ctx, &pb.Signup{
		First_Name: user.FirstName,
		Last_Name:  user.LastName,
		Email:      user.Email,
		Password:   user.Password,
		Phone:      user.Phone,
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
		"Message": "Go to verfication page & enter otp",
		"Data":    response,
	})
}

// VerificationHandler handles the verification process.
func VerificationHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var verificationDetails dto.OTP
	if err := c.BindJSON(&verificationDetails); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	response, err := client.VerifyUser(ctx, &pb.OTP{
		Email: verificationDetails.Email,
		Otp:   verificationDetails.OTP,
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
		"Message": "Signup request successful",
		"Data":    response,
	})

}

// UserLoginHandler function will send the login request to client.
func UserLoginHandler(c *gin.Context, client pb.UserServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	var user dto.Login
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error while binding json",
			"Error":   err.Error()})
		return
	}

	// if user.IsBlocked {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"status":  http.StatusBadRequest,
	// 		"Message": "user blocked by admin",
	// 	})
	// 	return
	// }

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "validation error",
			"Error":   err.Error()})
		return
	}

	response, err := client.UserLogin(ctx, &pb.Login{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest,
			"Message": "error in client response",
			"Error":   err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Status":  http.StatusAccepted,
		"Message": "Login successful",
		"Data":    response,
	})
}
