package school

import "github.com/axiaoxin/pink-lady/app/models"

type Grade struct {
	models.BaseModel
	Name      string
	GradeType GradeType
}
