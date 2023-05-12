package persist

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Client = &gorm.DB{}

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(127.0.0.1:3306)/xcelldb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//db.AutoMigrate()
	Client = db
	log.Println(":))")
}
