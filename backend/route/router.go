package route

import (
	"updownloader-backend/handler"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	updownRouter := g.Group("/updown")

	initUpdownRouter(updownRouter)

	g.Static("/webserver", "/data/updownloader")
	g.GET("/usage", handler.Usage)
}

func initUpdownRouter(r *gin.RouterGroup) {
	r.POST("/file", handler.AddFile)
	r.POST("/text", handler.AddText)
	r.GET("/all", handler.AllRecords)
	r.GET("/record/:code", handler.ReadRecord)
	r.DELETE("/record/:code", handler.DeleteRecord)
}
