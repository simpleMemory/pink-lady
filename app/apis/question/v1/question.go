package v1

import (
	"github.com/axiaoxin/pink-lady/app/domain"
	questionService "github.com/axiaoxin/pink-lady/app/services/question"
	"github.com/axiaoxin/pink-lady/app/services/retcode"
	"github.com/axiaoxin/pink-lady/app/utils/response"
	"github.com/gin-gonic/gin"
)

//获取题目
func GetQuestion(c *gin.Context) {
	query := domain.QuestionQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
		return
	}

	data, err := questionService.GetQuestion(&query)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}

	response.JSON(c, retcode.Success, data)
}

//保存题目和答案
func SaveQuestionAndAnswer(c *gin.Context) {
	body := domain.QuestionDomainBody{}
	res := make(map[string]interface{})

	if err := c.ShouldBind(body); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
		return
	}

	idQuestion, errQuestion := questionService.SaveQuesion(body.QueestionBody)
	res["questionId"] = idQuestion
	if errQuestion != nil {
		response.JSON500(c, retcode.Failure, res)
		return
	}

	answerIds, errAnswer := questionService.SaveAnswerBatch(body.AnswerBody)
	res["answerIds"] = answerIds
	if errAnswer != nil {
		response.JSON500(c, retcode.Failure, res)
		return
	}

	response.JSON(c, retcode.Success, res)
}

//删除题目
func DelQuestion(c *gin.Context) {
	query := domain.QuestionQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
		return
	}

	err := questionService.DelQuestion(query)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}

	response.JSON(c, retcode.Success, nil)
}

//获取题目列表
func GetQuestions(c *gin.Context) {
	query := domain.QuestionQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
		return
	}

	data, err := questionService.GetQuestions(&query)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}

	response.JSON(c, retcode.Success, data)
}
