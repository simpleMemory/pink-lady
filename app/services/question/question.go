package question

import (
	"github.com/axiaoxin/pink-lady/app/domain"
	"github.com/axiaoxin/pink-lady/app/models/question"
	"github.com/axiaoxin/pink-lady/app/utils"
)

func GetQuestion(query *domain.QuestionQuery) (*question.Question, error) {
	data := &question.Question{}
	err := utils.DB.Where(query).First(data).Error
	return data, err
}

func SaveQuestionAndAnswer(body *domain.QuestionDomainBody) (*domain.BaseDomain, error) {

	idQuestion, errQuestion := SaveQuesion(body.QueestionBody)

	if nil != errQuestion {
		return &domain.BaseDomain{
			Msg:     "",
			Success: false,
		}, errQuestion
	}

	idsAnswer, errAnswer := SaveAnswerBatch(body.AnswerBody)

	if nil != errAnswer {
		return &domain.BaseDomain{
			Msg:     "",
			Success: false,
			Data: &domain.QuesionDomainResponse{
				QuestionId: idQuestion,
			},
		}, errAnswer
	}

	return &domain.BaseDomain{
		Msg:     "",
		Success: true,
		Data: &domain.QuesionDomainResponse{
			QuestionId: idQuestion, AnswerIds: idsAnswer,
		},
	}, nil
}

func SaveQuesion(body *domain.QueestionBody) (string, error) {
	data := question.Question{}
	err := utils.DB.FirstOrCreate(&data, body).Error
	return data.ID, err
}

func SaveAnswer(body *domain.AnswerBody) (string, error) {
	data := question.Answer{}
	err := utils.DB.FirstOrCreate(&data, body).Error
	return data.ID, err
}

func SaveAnswerBatch(body []*domain.AnswerBody) ([]string, error) {
	ids := make([]string, 10)
	for _, ones := range body {
		id, err := SaveAnswer(ones)
		ids = append(ids, id)
		if err != nil {
			return ids, err
		}
	}
	return ids, nil
}

func GetQuestions(query *domain.QuestionQuery) (*domain.PageDomain, error) {

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
		scopedb = scopedb.Where(&question.Question{Body: query.ID})
		totalCount = false
	}

	totalCount = false

	data := []question.Question{}
	scopedb = scopedb.Find(&data)
	if totalCount {
		utils.DB.Model(&question.Question{}).Count(&count)
	} else {
		scopedb.Count(&count)
	}

	return &domain.PageDomain{
		Data:       data,
		Pagination: utils.Paginate(count, query.PageNum, query.PageSize),
	}, scopedb.Error
}

func DelQuestion(query domain.QuestionQuery) error {
	return utils.DB.Model(&question.Question{}).Where("id = ?", query.ID).Update("del_flag", 1).Error
}
