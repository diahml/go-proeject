package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConncectDatabase(){
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/go-project"))
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}