package database

import (
	"log"
	"updownloader-backend/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	dsn := config.DBLink()
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

func InsertRecord(record *Record) error {
	result := db.Create(record)
	return result.Error
}

func QueryRecordByCode(code string) (Record, bool) {
	record := Record{}
	res := db.Where("code = ?", code).First(&record).Error == nil
	return record, res
}

func DeleteRecordByCode(code string) {
	log.Println("delete record by code", code)
	db.Delete(&Record{}, "code = ?", code)
}

func GetAllRecords() []Record {
	var ret []Record
	db.Find(&ret)
	return ret
}
