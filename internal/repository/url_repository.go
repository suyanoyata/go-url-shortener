package repository

import (
	"fmt"
	"go-url-shortener/internal/model"
	"go-url-shortener/internal/types"

	"gorm.io/gorm"
)

type UrlRepository interface {
	Create(url string) (types.CreateShortenUrlResponse, error)
	FindUrlBySlug(slug string) (model.Url, error)
}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{
		db: db,
	}
}

func (repository *urlRepository) Create(url string) (types.CreateShortenUrlResponse, error) {
	urlModel := model.Url{
		Url: url,
	}

	if err := repository.db.Create(&urlModel).Error; err != nil {
		return types.CreateShortenUrlResponse{}, err
	}

	return types.CreateShortenUrlResponse{
		Url: fmt.Sprintf("http://localhost:1040/%s", urlModel.ShortURL),
	}, nil
}

func (repository *urlRepository) FindUrlBySlug(slug string) (model.Url, error) {
	var urlModel model.Url

	if err := repository.db.First(&urlModel).Where(&model.Url{
		ShortURL: slug,
	}).Error; err != nil {
		return model.Url{}, err
	}

	return urlModel, nil
}
