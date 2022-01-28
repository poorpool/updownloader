package handler

import (
	"updownloader-backend/database"
	"updownloader-backend/service"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "pong",
	})
}

func AddText(c *gin.Context) {
	text := c.PostForm("text")
	if l := len(text); l < 1 {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "Text cannot be empty",
		})
		return
	} else if l > 50000 {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "Text is so long",
		})
		return
	}
	res, err := service.SaveText(text)
	if err == nil {
		c.JSON(200, gin.H{
			"status": 0,
			"msg":    res,
		})
	} else {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    err.Error(),
		})
	}
}

func ReadRecord(c *gin.Context) {
	code := c.Param("code")
	record, exi := database.QueryRecordByCode(code)
	if !exi {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    "Not exist",
		})
	} else if record.Kind == 1 {
		c.JSON(200, gin.H{
			"status": 0,
			"kind":   1,
			"msg":    service.ReadText(code),
		})
	}
}

func AddFile(c *gin.Context) {

}
