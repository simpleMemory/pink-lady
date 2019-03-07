package v1

import (
	"github.com/axiaoxin/pink-lady/app/domain"
	bookService "github.com/axiaoxin/pink-lady/app/services/book"
	"github.com/axiaoxin/pink-lady/app/services/retcode"
	"github.com/axiaoxin/pink-lady/app/utils/response"
	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	query := domain.BookQuery{}

	if err := c.ShouldBind(&query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}

	book, err := bookService.GetBook(&query)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}

	response.JSON(c, retcode.Success, book)
}

func Save(c *gin.Context) {
	body := &domain.BookBody{}

	if err := c.ShouldBind(body); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
		return
	}

	id, err := bookService.SaveBook(body)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}

	response.JSON(c, retcode.Success, id)
}

func Gets(c *gin.Context) {
	query := &domain.BookQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}

	data, err := bookService.GetBooks(query)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}
	response.JSON(c, retcode.Success, data)
}

func Del(c *gin.Context) {
	query := &domain.BookQuery{}

	if err := c.ShouldBind(query); err != nil {
		response.JSON400(c, retcode.InvalidParams, err.Error())
	}

	err := bookService.DelBook(query)

	if err != nil {
		response.JSON(c, retcode.Failure, err.Error())
		return
	}
	response.JSON(c, retcode.Success, nil)
}
