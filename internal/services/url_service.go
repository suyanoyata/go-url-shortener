package service

import (
	"go-url-shortener/internal/model"
	"go-url-shortener/internal/repository"
	"go-url-shortener/internal/types"

	"gorm.io/gorm"
)

type UrlService interface {
	Create(url string) (types.CreateShortenUrlResponse, error)
	FindUrlBySlug(slug string) (model.Url, error)
}

type urlService struct {
	urlRepository repository.UrlRepository
}

func NewUrlService(db *gorm.DB) UrlService {
	urlRepository := repository.NewUrlRepository(db)

	return &urlService{
		urlRepository: urlRepository,
	}
}

func (service *urlService) Create(url string) (types.CreateShortenUrlResponse, error) {
	return service.urlRepository.Create(url)
}

func (service *urlService) FindUrlBySlug(slug string) (model.Url, error) {
	return service.urlRepository.FindUrlBySlug(slug)
}
