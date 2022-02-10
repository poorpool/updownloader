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
}
