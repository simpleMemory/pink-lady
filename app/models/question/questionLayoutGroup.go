package question

import "github.com/axiaoxin/pink-lady/app/models"

//题目布局类型分组

type QuestionLayoutGroup struct {
	models.BaseModel

	//分布名称
	Name string
	//说明
	Remark string
}
