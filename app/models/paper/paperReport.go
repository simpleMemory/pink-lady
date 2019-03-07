package paper

import (
	"github.com/axiaoxin/pink-lady/app/models"
)

type PaperReport struct {
	models.BaseModel
	ReportUser string
	Paper      Paper
	Score      int
}
