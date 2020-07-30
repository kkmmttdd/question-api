package question_repository

import (
	question "github.com/kkmmttdd/question-api/src/domain"
	mysql_repository "github.com/kkmmttdd/question-api/src/repository/question/mysql"
)

type QuestionRepositoryInterface interface {
	FetchAll() ([]*question.Question, error)
}

func NewQuestionRepository() QuestionRepositoryInterface {
	repo := mysql_repository.MysqlRepository{}
	return &repo
}
