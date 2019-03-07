package paper

import "github.com/axiaoxin/pink-lady/app/models"

type Paper struct {
	models.BaseModel
	Name     string
	Title    string
	Remark   string
	Template PaperTemplate
}
