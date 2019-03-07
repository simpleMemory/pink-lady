package school

import "github.com/axiaoxin/pink-lady/app/models"

type Student struct {
	models.BaseModel
	Num  string
	Name string
}
