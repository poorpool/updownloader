package main

import (
	"math/rand"
	"time"
	"updownloader-backend/database"
	"updownloader-backend/middleware"
	"updownloader-backend/route"
	"updownloader-backend/service"

	"github.com/gin-gonic/gin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	service.InitDir()

	database.InitDatabase()

	r := gin.Default()

	r.Use(middleware.Cors())

	route.InitRouter(r)

	r.Run("0.0.0.0:10370") // listen and serve on 0.0.0.0:8080
}
