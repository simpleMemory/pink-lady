package paper

import (
	"github.com/axiaoxin/pink-lady/app/models"
	"github.com/axiaoxin/pink-lady/app/models/question"
)

type PaperQuestion struct {
	models.BaseModel
	Paper        Paper
	Question     question.Question
	QuestionType question.QuestionType
}
