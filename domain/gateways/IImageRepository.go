package gateways

import (
	"RenewCMS/domain/image"
	"mime/multipart"
)

type IImageRepository interface {
	Create(file multipart.File, fileHeader multipart.FileHeader) (image.Image, error)
	Delete(id uint32) error
}
