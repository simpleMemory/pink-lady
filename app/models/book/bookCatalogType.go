package book

import "github.com/axiaoxin/pink-lady/app/models"

type BookCatalogType struct {
	models.BaseModel
	Name   string
	Remark string
}
