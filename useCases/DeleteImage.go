package useCases

import (
	"RenewCMS/domain/gateways"
)

type DeleteImageUseCase struct {
	imageRepository gateways.IImageRepository
}

func NewDeleteImageUseCase(imageRepository gateways.IImageRepository) *DeleteImageUseCase {
	return &DeleteImageUseCase{imageRepository}
}

func (g *DeleteImageUseCase) DeleteImage(imageId uint32) error {
	return g.imageRepository.Delete(imageId)
}
