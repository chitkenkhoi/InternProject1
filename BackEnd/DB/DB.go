package DB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func CreateDB() {
	dsn := "host=localhost user	=postgres password=teptom0792. dbname=test port=5432 sslmode=prefer connect_timeout=10 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	fmt.Print(DB)
}
