package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type database struct {
	orm *gorm.DB
}

var DB *database

func New() {
	// datasource の定義
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"cal_dev",
	)

	// connection 作成
	orm, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// このパッケージが持っている変数に代入
	DB = &database{orm: orm}
}
