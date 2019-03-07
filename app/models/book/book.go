package book

import (
	"github.com/axiaoxin/pink-lady/app/models"
)

type Book struct {
	models.BaseModel
	Name     string `gorm:"type:varchar(255);not null;default:''" json:"name" binding:"required,max=255"`
	Title    string `gorm:"type:varchar(255);not null;default:''" json:"title" binding:"max=255"`
	SubTitle string `gorm:"type:varchar(255);not null;default:''" json:"subTitle" binding:"max=255"`

	Version string `gorm:"type:varchar(255);not null;default:''" json:"subTitle" binding:"max=255"`
}

func (Book) TableName() string {
	return "book"
}

func init() {
	models.MigrationModels = append(models.MigrationModels, &Book{})
}
