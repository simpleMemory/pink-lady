package area

import "github.com/axiaoxin/pink-lady/app/models"

type AreaType struct {
	models.BaseModel
	Name  string
	Level string
}
