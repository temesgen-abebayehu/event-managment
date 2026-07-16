package cloudinary

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Client struct {
	cld *cloudinary.Cloudinary
}

func NewClient(cloudName, apiKey, apiSecret string) (*Client, error) {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}
	return &Client{cld: cld}, nil
}

type UploadResult struct {
	URL      string `json:"url"`
	PublicID string `json:"public_id"`
}

// Upload uploads a file with unique naming: events/{user_id}/{event_id}/{timestamp}_{uuid}
func (c *Client) Upload(file multipart.File, userID, eventID string) (*UploadResult, error) {
	ctx := context.Background()

	// Generate unique filename
	timestamp := time.Now().Unix()
	uniqueID := uuid.New().String()[:8]
	publicID := fmt.Sprintf("events/%s/%s/%d_%s", userID, eventID, timestamp, uniqueID)

	// Upload to Cloudinary
	uploadResult, err := c.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: publicID,
		Folder:   "events",
	})
	if err != nil {
		return nil, err
	}

	return &UploadResult{
		URL:      uploadResult.SecureURL,
		PublicID: uploadResult.PublicID,
	}, nil
}

// Delete deletes an image from Cloudinary
func (c *Client) Delete(publicID string) error {
	ctx := context.Background()
	_, err := c.cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: publicID,
	})
	return err
}

// DeleteMultiple deletes multiple images from Cloudinary
func (c *Client) DeleteMultiple(publicIDs []string) error {
	for _, publicID := range publicIDs {
		if err := c.Delete(publicID); err != nil {
			return err
		}
	}
	return nil
}
