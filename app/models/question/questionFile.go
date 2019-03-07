package question

import "github.com/axiaoxin/pink-lady/app/models"

//题目文件

type QuestionFile struct {
	models.BaseModel
	//所属题目
	Question Question
	//文件
	File string
	//文件名
	Name string
	//文件标签
	Tag string
	//排序
	Sort int
	//布局标示
	LayoutId string
}
