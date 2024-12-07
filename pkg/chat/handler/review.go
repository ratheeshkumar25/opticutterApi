package handler

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
// 	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/pb"
// )

// // SubmitReview handles review submission for a material.
// func (c *Chat) SubmitReview(ctx *gin.Context) {
// 	var req struct {
// 		ReviewText string `json:"review_text" binding:"required"`
// 		Rating     int32  `json:"rating" binding:"required"`
// 		MaterialID uint32 `json:"material_id" binding:"required"`
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

// 	// Call service for review submission
// 	_, err := c.client.SubmitReview(ctx, &pb.ReviewRequest{
// 		UserId:     uint32(userID.(int)),
// 		MaterialId: req.MaterialID,
// 		ReviewText: req.ReviewText,
// 		Rating:     req.Rating,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit review", "details": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "Review submitted successfully"})
// }

// // FetchReviews handles fetching reviews for a given material.
// func (c *Chat) FetchReviews(ctx *gin.Context) {
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

// 	// Call service to fetch reviews
// 	reviews, err := c.client.FetchReviews(ctx, &pb.MaterialID{Id: uint32(materialID)})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews", "details": err.Error()})
// 		return
// 	}

// 	// Map reviews to DTO
// 	var dtoReviews []dto.Review
// 	for _, review := range reviews.Reviews {
// 		dtoReviews = append(dtoReviews, dto.Review{
// 			ID:         review.ReviewId,
// 			UserID:     review.UserId,
// 			MaterialID: review.MaterialId,
// 			ReviewText: review.ReviewText,
// 			Rating:     review.Rating,
// 			Timestamp:  review.Timestamp,
// 		})
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"reviews": dtoReviews})
// }

// // UploadVideo handles video chunk upload for a material.
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

// // FetchVideos handles fetching uploaded videos for a given material.
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

// // SubmitReview handles review submission for a material.
// func SubmitReview(c *gin.Context, client pb.ChatServiceClient) {
// 	var req struct {
// 		UserID     uint32 `json:"user_id" binding:"required"`
// 		MaterialID uint32 `json:"material_id" binding:"required"`
// 		ReviewText string `json:"review_text" binding:"required"`
// 		Rating     int32  `json:"rating" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		log.Printf("Invalid request: %v", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
// 		return
// 	}
// 	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
// 	defer cancel()

// 	response, err := client.SubmitReview(ctx, &pb.ReviewRequest{
// 		UserId:     req.UserID,
// 		MaterialId: req.MaterialID,
// 		ReviewText: req.ReviewText,
// 		Rating:     req.Rating,
// 	})

// 	if err != nil {
// 		log.Printf("Failed to submit review: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit review", "details": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": response.Message})
// }

// // FetchReviews handles fetching reviews for a given material.
// func FetchReviews(c *gin.Context, client pb.ChatServiceClient) {
// 	materialIDStr := c.Param("material_id")
// 	materialID, err := strconv.Atoi(materialIDStr)
// 	if err != nil {
// 		log.Printf("Invalid material ID: %v", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
// 	defer cancel()

// 	response, err := client.FetchReviews(ctx, &pb.MaterialID{
// 		Id: uint32(materialID),
// 	})
// 	if err != nil {
// 		log.Printf("Failed to fetch reviews: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reviews", "details": err.Error()})
// 		return
// 	}

// 	// Construct the response
// 	reviews := make([]dto.Review, len(response.Reviews))
// 	for i, review := range response.Reviews {
// 		// Convert Timestamp string to time.Time
// 		timestamp, err := time.Parse(time.RFC3339, review.Timestamp)
// 		if err != nil {
// 			log.Printf("Invalid timestamp format: %v", err)
// 			timestamp = time.Time{}
// 		}
// 		reviews[i] = dto.Review{
// 			ID:         review.ReviewId,
// 			UserID:     review.UserId,
// 			MaterialID: review.MaterialId,
// 			ReviewText: review.ReviewText,
// 			Rating:     review.Rating,
// 			Timestamp:  timestamp,
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
// }

// // UploadVideo handles video chunk upload for a material.
// func UploadVideo(c *gin.Context, client pb.ChatServiceClient) {
// 	var req struct {
// 		VideoID      string `json:"video_id" binding:"required"`
// 		UserID       uint32 `json:"user_id" binding:"required"`
// 		MaterialID   uint32 `json:"material_id" binding:"required"`
// 		ChunkOrder   int32  `json:"chunk_order" binding:"required"`
// 		IsFirstChunk bool   `json:"is_first_chunk" binding:"required"`
// 		IsLastChunk  bool   `json:"is_last_chunk" binding:"required"`
// 		FileName     string `json:"file_name" binding:"required"`
// 		VideoURL     string `json:"video_url" binding:"required"`
// 		VideoData    []byte `json:"video_data" binding:"required"`
// 	}

// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		log.Printf("Invalid request: %v", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
// 	defer cancel()

// 	response, err := client.AddVideoChunk(ctx, &pb.VideoUploadRequest{
// 		VideoId:      req.VideoID,
// 		UserId:       req.UserID,
// 		MaterialId:   req.MaterialID,
// 		ChunkOrder:   req.ChunkOrder,
// 		IsFirstChunk: req.IsFirstChunk,
// 		IsLastChunk:  req.IsLastChunk,
// 		FileName:     req.FileName,
// 		VideoUrl:     req.VideoURL,
// 		VideoData:    req.VideoData,
// 	})
// 	if err != nil {
// 		log.Printf("Failed to upload video chunk: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video chunk", "details": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": response.Message, "video_id": response.VideoId})
// }

// // FetchVideos handles fetching uploaded videos for a given material.
// func FetchVideos(c *gin.Context, client pb.ChatServiceClient) {
// 	materialIDStr := c.Param("material_id")
// 	materialID, err := strconv.Atoi(materialIDStr)
// 	if err != nil {
// 		log.Printf("Invalid material ID: %v", err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
// 	defer cancel()

// 	response, err := client.FetchVideos(ctx, &pb.FetchVideoRequest{
// 		MaterialId: uint32(materialID),
// 	})
// 	if err != nil {
// 		log.Printf("Failed to fetch videos: %v", err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch videos", "details": err.Error()})
// 		return
// 	}

// 	// Construct the response
// 	videos := make([]dto.VideoMetadata, len(response.Videos))
// 	for i, video := range response.Videos {
// 		timestamp, err := time.Parse(time.RFC3339, video.Timestamp)
// 		if err != nil {
// 			log.Printf("Invalid timestamp format: %v", err)
// 			timestamp = time.Time{}
// 		}
// 		videos[i] = dto.VideoMetadata{
// 			VideoID:    video.VideoId,
// 			MaterialID: video.MaterialId,
// 			UserID:     video.UserId,
// 			FileName:   video.FileName,
// 			VideoURL:   video.VideoUrl,
// 			Timestamp:  timestamp,
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{"videos": videos})
// }
