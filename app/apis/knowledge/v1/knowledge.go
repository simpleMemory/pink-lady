package v1

import (
	"github.com/axiaoxin/pink-lady/app/domain"
	"github.com/axiaoxin/pink-lady/app/services/retcode"
	"github.com/axiaoxin/pink-lady/app/utils/response"
	"github.com/gin-gonic/gin"
)

func GetKnowledge(c *gin.Context) {
	query := domain.KnowledgeQuery{}

	if err := c.ShouldBind(&query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}
}

func SaveKnowledge(c *gin.Context) {
	data := domain.KnowledgeBody{}

	if err := c.ShouldBind(&data); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}

}

func DelKnowledge(c *gin.Context) {
	query := &domain.BookQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}
}

func GetKonwledges(c *gin.Context) {
	query := &domain.BookQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}
}
