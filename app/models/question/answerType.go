package question

import "github.com/axiaoxin/pink-lady/app/models"

//答案类型:文字？图片等。。
type AnswerType struct {
	models.BaseModel
	//类型名称
	Name string
	//类型标示
	TypeFlag string
}
