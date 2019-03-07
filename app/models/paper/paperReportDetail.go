package paper

import "github.com/axiaoxin/pink-lady/app/models"

type PaperReportDetail struct {
	models.BaseModel

	Paper         Paper
	PaperQuestion PaperQuestion
	Score         int
}
