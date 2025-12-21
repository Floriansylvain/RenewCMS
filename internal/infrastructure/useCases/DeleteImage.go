package useCases

import "RenewCMS/internal/domain/image"

type DeleteImageUseCase struct {
	imageRepository image.Repository
}

func NewDeleteImageUseCase(imageRepository image.Repository) *DeleteImageUseCase {
	return &DeleteImageUseCase{imageRepository}
}

func (g *DeleteImageUseCase) DeleteImage(imageId uint32) error {
	return g.imageRepository.Delete(imageId)
}
