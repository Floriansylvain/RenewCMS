package useCases

import (
	"RenewCMS/internal/domain/image"
	"mime/multipart"
)

type CreateImageUseCase struct {
	imageRepository image.Repository
}

func NewCreateImageUseCase(imageRepository image.Repository) *CreateImageUseCase {
	return &CreateImageUseCase{imageRepository}
}

func (g *CreateImageUseCase) CreateImage(file multipart.File, fileHeader multipart.FileHeader) (image.Image, error) {
	return g.imageRepository.Create(file, fileHeader)
}
