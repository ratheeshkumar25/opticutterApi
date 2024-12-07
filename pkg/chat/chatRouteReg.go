package chat

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/handler"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/pb"
)

// Chat handles the websocket connection in the chat page.
func (c *Chat) Chat(ctx *gin.Context) {
	handler.HandleWebSocketConnection(ctx, c.client, c.userClient)
}

// VideoCall handles initiating a video call.
func (c *Chat) VideoCall(ctx *gin.Context) {
	handler.StartVideoCall(ctx, c.client)
}

// ChatPage handles to load the chat page
func (c *Chat) ChatPage(ctx *gin.Context) {
	handler.ChatPage(ctx, c.client)
}

// SubmitReview handles review submission for a material.
func (c *Chat) SubmitReview(ctx *gin.Context) {
	var req struct {
		ReviewText string `json:"review_text" binding:"required"`
		Rating     int32  `json:"rating" binding:"required"`
		MaterialID uint32 `json:"material_id" binding:"required"`
	}

	// Retrieve user_id from context (set by middleware)
	userID, _ := ctx.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
		return
	}

	// Check if userID is of the correct type (uint), and then convert it to uint32
	uid, ok := userID.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}

	// Convert uint to uint32
	userID32 := uint32(uid)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	// Call service for review submission
	_, err := c.client.SubmitReview(ctx, &pb.ReviewRequest{
		UserId:     userID32,
		MaterialId: req.MaterialID,
		ReviewText: req.ReviewText,
		Rating:     req.Rating,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit review", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Review submitted successfully"})
}

// FetchReviews handles fetching reviews for a given material.
func (c *Chat) FetchReviews(ctx *gin.Context) {
	materialIDStr := ctx.Param("material_id")
	log.Println("materialstr", materialIDStr)
	materialID, err := strconv.Atoi(materialIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	// Retrieve user_id from context (set by middleware)
	userID, _ := ctx.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
		return
	}

	// Call service to fetch reviews
	reviews, err := c.client.FetchReviews(ctx, &pb.MaterialID{Id: uint32(materialID)})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews", "details": err.Error()})
		return
	}

	// Map reviews to DTO
	var dtoReviews []dto.Review
	for _, review := range reviews.Reviews {
		timestamp, err := time.Parse(time.RFC3339, review.Timestamp)
		if err != nil {
			log.Printf("Invalid timestamp format: %v", err)
			timestamp = time.Time{}
		}
		dtoReviews = append(dtoReviews, dto.Review{
			ID:         review.ReviewId,
			UserID:     review.UserId,
			MaterialID: review.MaterialId,
			ReviewText: review.ReviewText,
			Rating:     review.Rating,
			Timestamp:  timestamp,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"reviews": dtoReviews})
}

// UploadVideo handles video chunk upload for a material.
func (c *Chat) UploadVideo(ctx *gin.Context) {
	// Parse the multipart form
	if err := ctx.Request.ParseMultipartForm(32 << 20); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse multipart form", "details": err.Error()})
		return
	}

	// Retrieve the form fields
	videoID := ctx.Request.FormValue("video_id")
	materialID, err := strconv.Atoi(ctx.Request.FormValue("material_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID", "details": err.Error()})
		return
	}
	fileName := ctx.Request.FormValue("file_name")
	videoURL := ctx.Request.FormValue("video_url")
	isLastChunk := ctx.DefaultQuery("is_last_chunk", "false") == "true"

	// Retrieve the uploaded video file
	file, _, err := ctx.Request.FormFile("video_data")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No video file uploaded", "details": err.Error()})
		return
	}
	defer file.Close()

	// Read the video data into a byte slice
	videoBytes, err := io.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read video data", "details": err.Error()})
		return
	}

	// Determine whether this is the first chunk or not
	isFirstChunk := videoID == ""

	// Call the ChatService gRPC method
	response, err := c.client.AddVideoChunk(ctx, &pb.VideoUploadRequest{
		VideoId:      videoID,
		MaterialId:   uint32(materialID),
		FileName:     fileName,
		VideoUrl:     videoURL,
		VideoData:    videoBytes,
		IsFirstChunk: isFirstChunk,
		IsLastChunk:  isLastChunk,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video chunk", "details": err.Error()})
		return
	}

	// For the first chunk, return the new VideoID
	if isFirstChunk {
		ctx.JSON(http.StatusOK, gin.H{
			"message":  "First chunk uploaded successfully",
			"video_id": response.VideoId,
		})
		return
	}

	// For subsequent chunks, respond with success
	ctx.JSON(http.StatusOK, gin.H{"message": "Video chunk uploaded successfully"})
}

// FetchVideos handles fetching uploaded videos for a given material.
func (c *Chat) FetchVideos(ctx *gin.Context) {
	materialIDStr := ctx.Param("material_id")
	materialID, err := strconv.Atoi(materialIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	// Retrieve user_id from context (set by middleware)
	userID, _ := ctx.Get("user_id")
	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
		return
	}

	// Call service to fetch videos
	videos, err := c.client.FetchVideos(ctx, &pb.FetchVideoRequest{MaterialId: uint32(materialID)})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos", "details": err.Error()})
		return
	}

	// Map videos to DTO
	var dtoVideos []dto.VideoMetadata
	for _, video := range videos.Videos {
		dtoVideos = append(dtoVideos, dto.VideoMetadata{
			VideoID:    video.VideoId,
			MaterialID: video.MaterialId,
			FileName:   video.FileName,
			VideoURL:   video.VideoUrl,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"videos": dtoVideos})
}

// func (c *Chat) UploadVideo(ctx *gin.Context) {
// 	var req struct {
// 		VideoID    string `json:"video_id" binding:"required"`
// 		MaterialID uint32 `json:"material_id" binding:"required"`
// 		FileName   string `json:"file_name" binding:"required"`
// 		VideoURL   string `json:"video_url" binding:"required"`
// 		VideoData  []byte `json:"video_data" binding:"required"`
// 	}

// 	// Retrieve user_id from context (set by middleware)
// 	userID, _ := ctx.Get("user_id")
// 	if userID == nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
// 		return
// 	}

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
// 		return
// 	}

// 	// Call service to upload video chunk
// 	_, err := c.client.AddVideoChunk(ctx, &pb.VideoUploadRequest{
// 		VideoId:    req.VideoID,
// 		MaterialId: req.MaterialID,
// 		FileName:   req.FileName,
// 		VideoUrl:   req.VideoURL,
// 		VideoData:  req.VideoData,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video chunk", "details": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "Video uploaded successfully"})
// }

// FetchVideos handles fetching uploaded videos for a given material.
// func (c *Chat) FetchVideos(ctx *gin.Context) {
// 	materialIDStr := ctx.Param("material_id")
// 	materialID, err := strconv.Atoi(materialIDStr)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
// 		return
// 	}

// 	// Retrieve user_id from context (set by middleware)
// 	userID, _ := ctx.Get("user_id")
// 	if userID == nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User is not authenticated"})
// 		return
// 	}

// 	// Call service to fetch videos
// 	videos, err := c.client.FetchVideos(ctx, &pb.FetchVideoRequest{MaterialId: uint32(materialID)})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos", "details": err.Error()})
// 		return
// 	}

// 	// Map videos to DTO
// 	var dtoVideos []dto.VideoMetadata
// 	for _, video := range videos.Videos {
// 		dtoVideos = append(dtoVideos, dto.VideoMetadata{
// 			VideoID:    video.VideoId,
// 			MaterialID: video.MaterialId,
// 			FileName:   video.FileName,
// 			VideoURL:   video.VideoUrl,
// 		})
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"videos": dtoVideos})
// }
