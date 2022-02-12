package database

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Code       string // 获取文件的识别码
	Kind       int    // 1 for text, 2 for file
	Filename   string
	ExpireTime time.Time
	// path: basepath/Code/Filename
	// for text, filename is Code.txt
	// for file, filename is its original name
	// TODO: go是不是有文档注释啊
}
