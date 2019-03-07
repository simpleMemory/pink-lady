package paper

import "github.com/axiaoxin/pink-lady/app/models"

type PaperTemplate struct {
	models.BaseModel
	Name   string
	Remark string
}
