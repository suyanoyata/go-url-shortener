package controllers

import (
	service "go-url-shortener/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UrlController interface {
	Create(ctx *gin.Context)
	FindUrlBySlug(ctx *gin.Context)
}

type urlController struct {
	service service.UrlService
	db      *gorm.DB
}

func NewUrlController(db *gorm.DB) UrlController {
	urlService := service.NewUrlService(db)

	return &urlController{
		service: urlService,
		db:      db,
	}
}

type ShortenUrlJsonBody struct {
	Url string `json:"url" binding:"required"`
}

func (controller *urlController) Create(ctx *gin.Context) {
	var body ShortenUrlJsonBody

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request: URL is required",
			"details": err.Error(),
		})

		return
	}

	data, err := controller.service.Create(body.Url)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"data": data,
		},
	)
}

func (controller *urlController) FindUrlBySlug(ctx *gin.Context) {
	slug := ctx.Param("slug")

	data, err := controller.service.FindUrlBySlug(slug)

	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.Redirect(301, data.Url)
}
