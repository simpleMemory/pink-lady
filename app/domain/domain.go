package domain

type BaseDomain struct {
	Code    int
	Success bool
	Msg     string
	Data    interface{}
}

type BaseQuery struct {
	ID string `form:"id"`
	// page number
	PageNum int `form:"pageNum"`
	// page size
	PageSize int `form:"pageSize"`
	// how to order
	Order string `form:"order"`
}
