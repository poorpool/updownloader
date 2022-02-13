package main

import (
	"math/rand"
	"time"
	"updownloader-backend/crontab"
	"updownloader-backend/database"
	"updownloader-backend/middleware"
	"updownloader-backend/route"
	"updownloader-backend/service"

	// limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	service.InitDir()

	database.InitDatabase()

	crontab.InitCron()

	r := gin.Default()

	// r.MaxMultipartMemory = 500 << 20 // 500 MiB 这个并不是限制文件大小的
	// r.Use(limits.RequestSizeLimiter(500 << 20)) 这个限制的比较野蛮粗暴还会有异常

	r.Use(middleware.Cors())

	route.InitRouter(r)

	r.Run("0.0.0.0:10370") // listen and serve on 0.0.0.0:8080
}
