package question

import "github.com/axiaoxin/pink-lady/app/models"

//题目布局

type QuestionLayout struct {
	models.BaseModel

	//布局名称
	Name string
	//布局模板
	LayoutTemplate LayoutTemplate

	//布局预览图
	PreviewImg string
	//布局预览缩略图
	Thumbnail string
}
