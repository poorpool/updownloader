package database

import (
	"fmt"
	"testing"
	"time"
)

func Test_SaveRecord(t *testing.T) {
	InitDatabase()
	record := Record{
		Code:       "edva",
		Kind:       1,
		Filename:   "edva.txt",
		ExpireTime: time.Now(),
	}
	err := InsertRecord(record)
	if err == nil {
		t.Log("Save Record Passed")
	} else {
		t.Error(err)
	}
}

func Test_GetRecord(t *testing.T) {
	InitDatabase()
	record, res := QueryRecordByCode("edva")
	fmt.Print(record, res)
	record, res = QueryRecordByCode("ddddddd")
	fmt.Print(record, res)
}
