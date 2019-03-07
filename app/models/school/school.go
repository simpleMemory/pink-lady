package school

import (
	"github.com/axiaoxin/pink-lady/app/models"
	"github.com/axiaoxin/pink-lady/app/models/area"
)

type School struct {
	models.BaseModel
	Name string
	Area area.Area
}
