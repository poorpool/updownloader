package route

import (
	"updownloader-backend/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	updownRouter := g.Group("/updown")
	// rg := g.Group("/admin")
	initUpdownRouter(updownRouter)
}

func initUpdownRouter(r *gin.RouterGroup) {
	r.GET("/", handler.Ping)
	r.POST("/file", handler.AddFile)
	r.POST("/text", handler.AddText)
	r.GET("/record/:code", handler.ReadRecord)
}
