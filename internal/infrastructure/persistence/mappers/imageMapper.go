package mappers

import (
	"RenewCMS/internal/domain/image"
	"RenewCMS/internal/infrastructure/persistence/models"
)

func ImageToDomain(imageModel models.Image) image.Image {
	return image.Image{
		ID:        imageModel.ID,
		Path:      imageModel.Path,
		ArticleID: imageModel.ArticleID,
		CreatedAt: imageModel.CreatedAt,
		UpdatedAt: imageModel.UpdatedAt,
	}
}

func ImageToModel(imageDomain image.Image) models.Image {
	return models.Image{
		Path:      imageDomain.Path,
		ArticleID: imageDomain.ArticleID,
		CreatedAt: imageDomain.CreatedAt,
		UpdatedAt: imageDomain.UpdatedAt,
	}
}
