package useCases

import (
	"RenewCMS/domain/gateways"
)

type GetPageUseCase struct {
	pageRepository gateways.IPageRepository
}

func NewGetPageUseCase(pageRepository gateways.IPageRepository) *GetPageUseCase {
	return &GetPageUseCase{pageRepository}
}

func (g *GetPageUseCase) GetPage(name string, data any) ([]byte, error) {
	return g.pageRepository.Get(name, data)
}
