	package dto

	import "time"

	// Review represents a review submitted by a user for a material.
	type Review struct {
		ID         uint32    `json:"review_id"`
		UserID     uint32    `json:"user_id"`
		MaterialID uint32    `json:"material_id"`
		ReviewText string    `json:"review_text"`
		Rating     int32     `json:"rating"`
		Timestamp  time.Time `json:"timestamp"`
	}

	// VideoMetadata represents the metadata for a video uploaded.
	type VideoMetadata struct {
		VideoID    string    `json:"video_id"`
		MaterialID uint32    `json:"material_id"`
		UserID     uint32    `json:"user_id"`
		FileName   string    `json:"file_name"`
		VideoURL   string    `json:"video_url"`
		Timestamp  time.Time `json:"timestamp"`
	}

	// VideoChunk represents a chunk of the video being uploaded.
	type VideoChunk struct {
		ChunkID    string `json:"chunk_id"`
		ChunkData  []byte `json:"chunk_data"`
		ChunkOrder int    `json:"chunk_order"`
	}
