package image

type Repository interface {
	Create(input ImageInput) (Image, error)
	Delete(id uint32) error
}
