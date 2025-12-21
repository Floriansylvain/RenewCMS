package persistence

import (
	domain "RenewCMS/internal/domain/image"
	"RenewCMS/internal/infrastructure/persistence/mappers"
	entity "RenewCMS/internal/infrastructure/persistence/models"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{db}
}

var contentTypeExtensions = map[string]string{
	"image/png":     ".png",
	"image/jpeg":    ".jpeg",
	"image/webp":    ".webp",
	"image/svg+xml": ".svg",
}

func (i ImageRepository) Create(file multipart.File, fileHeader multipart.FileHeader) (domain.Image, error) {
	uploadDir := os.Getenv("UPLOAD_DIR")

	fileBytes := make([]byte, fileHeader.Size)
	_, err := file.Read(fileBytes)
	if err != nil {
		return domain.Image{}, err
	}

	err = os.MkdirAll(uploadDir, 0755)
	if err != nil {
		return domain.Image{}, err
	}

	contentType := fileHeader.Header.Get("Content-Type")
	extension := contentTypeExtensions[contentType]
	if extension == "" {
		return domain.Image{}, errors.New("the file must be a PNG, JPEG, WEBP, or SVG image")
	}

	newName := uuid.NewString() + extension
	finalPath := filepath.Join(uploadDir, newName)
	err = os.WriteFile(finalPath, fileBytes, 0666)
	if err != nil {
		return domain.Image{}, err
	}

	newImage := i.db.Create(&domain.Image{Path: "/static/uploadedImages/" + newName})
	var createdImage entity.Image
	newImage.Scan(&createdImage)

	return mappers.ImageToDomain(createdImage), nil
}

func (i ImageRepository) Delete(id uint32) error {
	uploadDir := os.Getenv("UPLOAD_DIR")
	var image entity.Image
	err := i.db.Model(&entity.Image{}).First(&image, id).Error
	if err != nil {
		return err
	}

	fileName := filepath.Base(image.Path)
	err = os.Remove(filepath.Join(uploadDir, fileName))
	if err != nil {
		return err
	}

	return i.db.Delete(&entity.Image{}, id).Error
}

var _ domain.Repository = &ImageRepository{}
