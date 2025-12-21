package image

import (
	"mime/multipart"
)

type Repository interface {
	Create(file multipart.File, fileHeader multipart.FileHeader) (Image, error)
	Delete(id uint32) error
}
