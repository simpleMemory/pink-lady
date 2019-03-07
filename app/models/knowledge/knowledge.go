package knowledge

import "github.com/axiaoxin/pink-lady/app/models"

//知识点

type Knowledge struct {
	models.BaseModel
	Title       string
	Description string
	Body        string
}
