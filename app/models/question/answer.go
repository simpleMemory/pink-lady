package question

import "github.com/axiaoxin/pink-lady/app/models"

//答案
type Answer struct {
	models.BaseModel
	//答案
	Answer string
	//答案类型
	AnswerType AnswerType
	//所属题目
	Question Question
	//是否正确项
	Proper int
}
