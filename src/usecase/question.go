package question_usecase

import (
	question "github.com/kkmmttdd/question-api/src/domain"
	question_repository "github.com/kkmmttdd/question-api/src/repository/question"
)


var question_repo question_repository.QuestionRepositoryInterface

func init () {
	question_repo = question_repository.NewQuestionRepository()
}

func NewQuestionUseCase() (QuestionUseCaseInterface, error) {
	return &QuestionUseCase{}, nil
}

type QuestionUseCaseInterface interface {
	FetchAll() ([]*question.Question, error)
}

type QuestionUseCase struct {}

func (usecase *QuestionUseCase) FetchAll() ([]*question.Question, error) {
	questions, _ := question_repo.FetchAll()
	return questions, nil
}
