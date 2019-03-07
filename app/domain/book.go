package domain

type BookBody struct {
	ID       string `form:"id"`
	Name     string `form:"name" binding:"required"`
	Title    string `form:"title" binding:"required"`
	SubTitle string `form:"subTitle" binding:"required"`
	Version  string `form:"version" binding:"required"`
}

type BookQuery struct {
	ID       string   `form:"id"`
	IDs      []string `form:"ids"`
	Name     string   `form:"name"`
	Title    string   `form:"title"`
	SubTitle string   `form:"subTitle"`
	Version  string   `form:"version"`
	BaseQuery
}
