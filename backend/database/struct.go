package database

import "time"

type Record struct {
	Code       string // 获取文件的识别码
	Kind       int    // 1 for text, 2 for file
	Filename   string
	ExpireTime time.Time
	// path: basepath/Code/Filename
}
