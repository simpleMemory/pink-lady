package question

import "github.com/axiaoxin/pink-lady/app/models"

//题目

type Question struct {
	models.BaseModel

	//题目类型
	QuestionType QuestionType
	//题目布局
	QuestionLayout QuestionLayout
	//题目布局-
	QuestionLayoutString string

	//题干内容
	Body string

	//答案
	Answers []Answer
	//公式
	//Formula string
	//附件(图片)
	QuestionFiles []QuestionFile
}
