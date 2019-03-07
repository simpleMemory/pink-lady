package book

import "github.com/axiaoxin/pink-lady/app/models"

//教材结构定义 http://staruml.io/download/releases/StarUML-3.0.2-x86_64.AppImage

type BookCatalog struct {
	models.BaseModel
	Name     string
	Title    string
	SubTitle string
	OrderNum int
	Type     BookCatalogType

	Parent BookCatalog `gorm:"column:pid"`
	Book   Book        `gorm:"column:book_Id"`
}
