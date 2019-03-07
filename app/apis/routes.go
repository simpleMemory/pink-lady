// Package apis routes.go provides register handlers on url
package apis

import (
	book "github.com/axiaoxin/pink-lady/app/apis/book/v1"
	question "github.com/axiaoxin/pink-lady/app/apis/question/v1"

	// need by swag
	_ "github.com/axiaoxin/pink-lady/app/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// RegisterRoutes add handlers on urls at there
func RegisterRoutes(app *gin.Engine) {

	// group x registered pink-lady default api
	x := app.Group("/x")
	{
		x.GET("/apidocs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		x.GET("/ping", Ping)
	}

	// group book api version 1
	bookGroup := app.Group("/book/v1")
	{
		bookGroup.GET("/get", book.Get)
		bookGroup.POST("/save", book.Save)
		bookGroup.POST("/list", book.Gets)
		bookGroup.POST("/delete", book.Del)
	}

	questionGroup := app.Group("/question/v1")
	{
		questionGroup.GET("/get", question.GetQuestion)
	}
}
