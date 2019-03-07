package domain

//just for question
type QueestionBody struct {
}

type QuestionQuery struct {
	BaseQuery
}

type AnswerBody struct {
}

//for question 、answer struct
type QuestionDomainBody struct {
	QueestionBody *QueestionBody
	AnswerBody    []*AnswerBody
}

type QuesionDomainResponse struct {
	QuestionId string
	AnswerIds  []string
}
