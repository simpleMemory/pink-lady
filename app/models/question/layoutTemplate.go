package question

import "github.com/axiaoxin/pink-lady/app/models"

//布局模板

type LayoutTemplate struct {
	models.BaseModel

	//模板名称
	Name string
	//HTML
	HtmlString string
	//说明
	Remark string
}
