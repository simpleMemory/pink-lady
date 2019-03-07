package school

import "github.com/axiaoxin/pink-lady/app/models"

type GradeType struct {
	models.BaseModel
	Name  string
	Level int
}
