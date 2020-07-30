package question

type Question struct {
	QuestionText string `gorm:"column:question_text"`
}