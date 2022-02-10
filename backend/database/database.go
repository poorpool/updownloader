package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	// TODO: add config.yaml
	dsn := "root:123456@tcp(127.0.0.1:3306)/updownloader?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Record{}) {
		db.AutoMigrate(&Record{})
	} else {
		db.Migrator().CreateTable(&Record{})
	}

	log.Println("Connected to Mysql!")
}

// TODO: 这里要用指针吗
func InsertRecord(record Record) error {
	result := db.Create(&record)
	return result.Error
}

func QueryRecordByCode(code string) (Record, bool) {
	record := Record{}
	res := db.Where("code = ?", code).First(&record).Error == nil
	return record, res
}
