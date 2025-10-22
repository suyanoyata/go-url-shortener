package model

import (
	"github.com/lucsky/cuid"
	"gorm.io/gorm"
)

type Url struct {
	ID       string `gorm:"primaryKey" json:"id"`
	Url      string `json:"url"`
	ShortURL string `gorm:"unique" json:"short_url"`
}

func (u *Url) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = cuid.New()
	u.ShortURL = cuid.Slug()

	return
}
