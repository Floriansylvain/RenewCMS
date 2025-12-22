package useCases

import (
	"RenewCMS/internal/domain/image"
)

type CreateImageUseCase struct {
	imageRepository image.Repository
}

func NewCreateImageUseCase(imageRepository image.Repository) *CreateImageUseCase {
	return &CreateImageUseCase{imageRepository}
}

func (g *CreateImageUseCase) CreateImage(input image.ImageInput) (image.Image, error) {
	return g.imageRepository.Create(input)
}
