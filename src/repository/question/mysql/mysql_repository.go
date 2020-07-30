package mysql_repository

import (
	"fmt"
	question "github.com/kkmmttdd/question-api/src/domain"
)
import "github.com/jinzhu/gorm"

var db *gorm.DB

func init () {
	var _ error
	var err error
	db, err = init_db()
	if err != nil {
		fmt.Println(err)
	}
	db.LogMode(true)
}

type MysqlRepository struct {
}

func (repo *MysqlRepository) FetchAll () ([]*question.Question, error){
	//qs := question.Question{QuestionText: "weeeeeeeeeeeeeeee"}
	//
	//qslist := []*question.Question{&qs}
	questions := []question.Question{}
	db.Find(&questions)
	for _, q := range questions {
		fmt.Println(q)
	}
	return []*question.Question{&questions[0]}, nil
}
