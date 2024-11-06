package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Helper to re-compute hash from email and userID
func hashPayload(email string, userID uint) string {
	h := sha256.New()
	h.Write([]byte(email + fmt.Sprint(userID)))
	return hex.EncodeToString(h.Sum(nil))
}

// Authorization checks the token is genuine or not.
func Authorization(key string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Token not found in header",
				"Data":    "",
				"Error":   "null token"})
			ctx.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		// Decode and parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(key), nil

		})

		fmt.Println("token rec", token)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Token not valid",
				"Data":    "",
				"Error":   err.Error()})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		fmt.Println("afrer claims", claims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Invalid token claims",
				"Data":    "",
				"Error":   ok})
			ctx.Abort()
			return
		}
		email, ok := claims["Email"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Email not found in claims",
				"Data":    "",
				"Error":   ok})
			ctx.Abort()
			return
		}

		userIDF, ok := claims["UserID"].(float64)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "UserID not found in token",
				"Data":    userIDF,
				"Error":   ok})
			ctx.Abort()
			return
		}

		// Verify payload hash
		payloadHash, ok := claims["PayloadHash"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Failed",
				"Message": "PayloadHash not found in claims",
				"Data":    "",
				"Error":   "PayloadHash not found",
			})
			ctx.Abort()
			return
		}

		if hashPayload(email, uint(userIDF)) != payloadHash {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Failed",
				"Message": "Token payload tampered",
				"Data":    "",
				"Error":   "Hash mismatch",
			})
			ctx.Abort()
			return
		}
		userID := uint(userIDF)
		ctx.Set("email", email)
		ctx.Set("user_id", userID)
		ctx.Next()
	}
}

func AdminhashPayload(email string) string {
	h := sha256.New()
	h.Write([]byte(email))
	return hex.EncodeToString(h.Sum(nil))
}

// AdminAuthorization checks the token is genuine or not.
func AdminAuthorization(key, role string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Token not found in header",
				"Data":    "",
				"Error":   "null token"})
			ctx.Abort()
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		// Decode and parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// Validate the algorithm
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(key), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Token not valid",
				"Data":    "",
				"Error":   err.Error()})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Invalid token claims",
				"Data":    "",
				"Error":   ok})
			ctx.Abort()
			return
		}
		email, ok := claims["Email"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "Email not found in claims",
				"Data":    "",
				"Error":   ok})
			ctx.Abort()
			return
		}

		claimRole, ok := claims["Role"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "role not found in token",
				"Data":    role,
				"Error":   ok})
			ctx.Abort()
			return
		}

		if role != claimRole {
			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
				"Message": "role not matching",
				"Data":    role,
				"Error":   ok})
			ctx.Abort()
			return
		}
		// Verify payload hash
		payloadHash, ok := claims["PayloadHash"].(string)
		if !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Failed",
				"Message": "PayloadHash not found in claims",
				"Data":    "",
				"Error":   "claims assertion failed",
			})
			ctx.Abort()
			return
		}

		if AdminhashPayload(email) != payloadHash {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Status":  "Failed",
				"Message": "Token payload tampered",
				"Data":    "",
				"Error":   "Hash mismatch",
			})
			ctx.Abort()
			return
		}

		ctx.Set("email", email)
		ctx.Set("role", role)
		ctx.Next()
	}
}

// // Authorization middleware validates user access token.
// func Authorization(key string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		tokenString := ctx.GetHeader("Authorization")

// 		if tokenString == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Token not found in header",
// 				"Data":    "",
// 				"Error":   "null token"})
// 			ctx.Abort()
// 			return
// 		}

// 		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

// 		// Decode and parse the token with additional validation
// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 			}
// 			return []byte(key), nil
// 		})

// 		if err != nil || !token.Valid {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Token not valid",
// 				"Data":    "",
// 				"Error":   err.Error()})
// 			ctx.Abort()
// 			return
// 		}

// 		// Check token claims
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Invalid token claims",
// 				"Data":    "",
// 				"Error":   "claims not found"})
// 			ctx.Abort()
// 			return
// 		}

// 		email, ok := claims[key].(string)
// 		log.Print("claimed email", email, ok)
// 		if !ok || email == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Email not found in token claims",
// 				"Error":   "missing email claim"})
// 			ctx.Abort()
// 			return
// 		}

// 		userIDF, ok := claims["UserID"].(float64)
// 		log.Print("claimied userID", userIDF)
// 		if !ok {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "UserID not found in token",
// 				"Data":    "",
// 				"Error":   "missing userID claim"})
// 			ctx.Abort()
// 			return
// 		}

// 		// Convert float64 userID to uint
// 		userID := uint(userIDF)

// 		// Set context values and pass to the next handler
// 		ctx.Set("email", email)
// 		ctx.Set("user_id", userID)
// 		ctx.Next()
// 	}
// }

// // AdminAuthorization validates admin access token.
// func AdminAuthorization(key, requiredRole string) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		tokenString := ctx.GetHeader("Authorization")

// 		if tokenString == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Token not found in header",
// 				"Data":    "",
// 				"Error":   "null token"})
// 			ctx.Abort()
// 			return
// 		}

// 		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

// 		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 			}
// 			return []byte(key), nil
// 		})

// 		if err != nil || !token.Valid {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Token not valid",
// 				"Data":    "",
// 				"Error":   err.Error()})
// 			ctx.Abort()
// 			return
// 		}

// 		// Validate claims
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Invalid token claims",
// 				"Error":   "claims not found"})
// 			ctx.Abort()
// 			return
// 		}

// 		email, ok := claims["Email"].(string)
// 		if !ok || email == "" {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message": "Email not found in token claims",
// 				"Error":   "missing email claim"})
// 			ctx.Abort()
// 			return
// 		}

// 		claimRole, ok := claims["Role"].(string)
// 		if !ok || claimRole != requiredRole {
// 			ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "Failed",
// 				"Message":      "Role not matching or not found",
// 				"ExpectedRole": requiredRole,
// 				"Error":        "role mismatch"})
// 			ctx.Abort()
// 			return
// 		}

// 		// Set context values for subsequent handlers
// 		ctx.Set("email", email)
// 		ctx.Set("role", claimRole)
// 		ctx.Next()
// 	}
// }
