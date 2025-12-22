package test

import (
	"RenewCMS/internal/domain/image"
	"RenewCMS/internal/infrastructure/useCases"
	"strings"
	"testing"
)

type mockImageRepo struct {
	images map[uint32]image.Image
}

func (m *mockImageRepo) Create(input image.ImageInput) (image.Image, error) {
	newId := uint32(len(m.images) + 1)
	img := image.Image{
		ID:   newId,
		Path: "/static/" + input.Filename,
	}
	m.images[newId] = img
	return img, nil
}

func (m *mockImageRepo) Delete(id uint32) error {
	delete(m.images, id)
	return nil
}

var _ image.Repository = &mockImageRepo{}

func TestImageUseCases(t *testing.T) {
	repo := &mockImageRepo{images: make(map[uint32]image.Image)}
	createUc := useCases.NewCreateImageUseCase(repo)
	deleteUc := useCases.NewDeleteImageUseCase(repo)

	t.Run("Create Image", func(t *testing.T) {
		input := image.ImageInput{
			Content:     strings.NewReader("fake-image-data"),
			Filename:    "test.png",
			ContentType: "image/png",
		}

		res, err := createUc.CreateImage(input)
		if err != nil {
			t.Fatalf("CreateImage failed: %v", err)
		}
		if res.ID == 0 {
			t.Errorf("Expected valid ID, got 0")
		}
		if !strings.Contains(res.Path, "test.png") {
			t.Errorf("Path mismatch: %s", res.Path)
		}
	})

	t.Run("Delete Image", func(t *testing.T) {
		repo.images[99] = image.Image{ID: 99, Path: "/static/delete-me.png"}

		err := deleteUc.DeleteImage(99)
		if err != nil {
			t.Fatalf("DeleteImage failed: %v", err)
		}

		if _, exists := repo.images[99]; exists {
			t.Errorf("Image 99 should have been deleted")
		}
	})
}
