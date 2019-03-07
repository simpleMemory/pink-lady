package area

import "github.com/axiaoxin/pink-lady/app/models"

//区域

type Area struct {
	models.BaseModel
	Name      string
	FullName  string
	ShortName string
	Code      string

	Parent Area
	Type   AreaType
}
