package book

import (
	"github.com/axiaoxin/pink-lady/app/domain"
	"github.com/axiaoxin/pink-lady/app/models/book"
	bookModels "github.com/axiaoxin/pink-lady/app/models/book"
	"github.com/axiaoxin/pink-lady/app/utils"
)

func GetBook(query *domain.BookQuery) (*book.Book, error) {
	data := &book.Book{}
	err := utils.DB.Where(query).First(data).Error
	return data, err
}

func SaveBook(body *domain.BookBody) (string, error) {
	data := bookModels.Book{}
	err := utils.DB.FirstOrCreate(&data, body).Error
	return data.ID, err
}

func GetBooks(query *domain.BookQuery) (*domain.PageDomain, error) {

	if query.Order == "" {
		query.Order = "id asc"
	}
	if query.PageNum <= 0 {
		query.PageNum = 1
	}

	offset := (query.PageNum - 1) * query.PageSize
	limit := query.PageSize
	if query.PageSize <= 0 {
		query.PageSize = -1
		offset = -1
		limit = -1
	}

	count := 0
	totalCount := true
	scopedb := utils.DB.Order(query.Order).Offset(offset).Limit(limit)

	if query.ID != "" {
		scopedb = scopedb.Where(&bookModels.Book{Name: query.ID})
		totalCount = false
	}
	if query.Name != "" {
		scopedb = scopedb.Where(&bookModels.Book{Name: query.Name})
		totalCount = false
	}
	if query.Title != "" {
		scopedb = scopedb.Where(&bookModels.Book{Title: query.Title})
		totalCount = false
	}
	if query.SubTitle != "" {
		scopedb = scopedb.Where(&bookModels.Book{SubTitle: query.SubTitle})
		totalCount = false
	}
	if query.Version != "" {
		scopedb = scopedb.Where(&bookModels.Book{Version: query.Version})
		totalCount = false
	}

	totalCount = false

	data := []bookModels.Book{}
	scopedb = scopedb.Find(&data)
	if totalCount {
		utils.DB.Model(&bookModels.Book{}).Count(&count)
	} else {
		scopedb.Count(&count)
	}

	return &domain.PageDomain{
		Data:       data,
		Pagination: utils.Paginate(count, query.PageNum, query.PageSize),
	}, scopedb.Error
}

func DelBook(query *domain.BookQuery) error {
	return utils.DB.Model(&bookModels.Book{}).Where("id = ?", query.ID).Update("del_flag", 1).Error
}
