package gateways

import (
	entity "RenewCMS/adapters/secondary/gateways/models"
	"RenewCMS/domain/gateways"
	domain "RenewCMS/domain/image"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mime/multipart"
	"os"
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

func mapImageToDomain(image entity.Image) domain.Image {
	return domain.FromDB(
		image.ID,
		image.Path,
		image.ArticleID,
		image.CreatedAt,
		image.UpdatedAt,
	)
}

func (i ImageRepository) Create(file multipart.File, fileHeader multipart.FileHeader) (domain.Image, error) {
	fileBytes := make([]byte, fileHeader.Size)
	_, err := file.Read(fileBytes)
	if err != nil {
		return domain.Image{}, err
	}

	err = os.MkdirAll("api/static/uploadedImages/", 0666)
	if err != nil {
		return domain.Image{}, err
	}

	contentType := fileHeader.Header.Get("Content-Type")
	extension := contentTypeExtensions[contentType]
	if extension == "" {
		return domain.Image{}, errors.New("the file must be a PNG, JPEG, WEBP, or SVG image")
	}

	newName := uuid.NewString() + extension
	err = os.WriteFile("api/static/uploadedImages/"+newName, fileBytes, 0666)
	if err != nil {
		return domain.Image{}, err
	}

	newImage := i.db.Create(&domain.Image{Path: "/static/uploadedImages/" + newName})
	var createdImage entity.Image
	newImage.Scan(&createdImage)

	return mapImageToDomain(createdImage), nil
}

func (i ImageRepository) Delete(id uint32) error {
	var image entity.Image
	err := i.db.Model(&entity.Image{}).First(&image, id).Error
	if err != nil {
		return err
	}

	err = os.Remove("api" + image.Path)
	if err != nil {
		return err
	}

	return i.db.Delete(&entity.Image{}, id).Error
}

var _ gateways.IImageRepository = &ImageRepository{}
