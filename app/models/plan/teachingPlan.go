package plan

import (
	"github.com/axiaoxin/pink-lady/app/models"
	"github.com/axiaoxin/pink-lady/app/models/book"
)

//教案

type TeachingPlan struct {
	models.BaseModel

	Book book.Book

	Name string
}
