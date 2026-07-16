package usecase

import (
	"errors"
	"event-management-backend/domain"
	"event-management-backend/infrastructure/cloudinary"
	"mime/multipart"
)

type UploadUsecase struct {
	cloudinary *cloudinary.Client
}

func NewUploadUsecase(cloudinary *cloudinary.Client) *UploadUsecase {
	return &UploadUsecase{cloudinary: cloudinary}
}

func (u *UploadUsecase) UploadFiles(files []*multipart.FileHeader, userID, eventID string) (*domain.UploadResponse, error) {
	if len(files) == 0 {
		return nil, errors.New("no files provided")
	}

	if len(files) > 10 {
		return nil, errors.New("maximum 10 files allowed")
	}

	var uploadedFiles []domain.UploadedFile

	for _, fileHeader := range files {
		// Validate file size (5MB max)
		if fileHeader.Size > 5*1024*1024 {
			return nil, errors.New("file size exceeds 5MB limit")
		}

		// Open file
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		// Upload to Cloudinary
		result, err := u.cloudinary.Upload(file, userID, eventID)
		if err != nil {
			return nil, err
		}

		uploadedFiles = append(uploadedFiles, domain.UploadedFile{
			URL:      result.URL,
			PublicID: result.PublicID,
		})
	}

	return &domain.UploadResponse{
		Files: uploadedFiles,
	}, nil
}

func (u *UploadUsecase) DeleteFiles(publicIDs []string) error {
	return u.cloudinary.DeleteMultiple(publicIDs)
}
