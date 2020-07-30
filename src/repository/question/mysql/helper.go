package mysql_repository
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init_db() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3307)/development")
	if err != nil {
		panic(err)
	}
	return db, nil
}