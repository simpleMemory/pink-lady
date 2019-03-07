package question

import "github.com/axiaoxin/pink-lady/app/models"

//题目类型

type QuestionType struct {
	models.BaseModel
	//类型名称
	Name string
	//说明
	Remark string
	//启用标记
	Flag int
}
