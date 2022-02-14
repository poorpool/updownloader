package handler

import (
	"net/http"
	"updownloader-backend/database"
	"updownloader-backend/service"

	"github.com/gin-gonic/gin"
)

func AddText(c *gin.Context) {
	var req AddTextReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(200, gin.H{
			"status": 1,
			"msg":    err,
		})
		return
	}
	if l := len(req.Text); l < 1 {
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
	res, err := service.SaveText(req.Text)
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
	} else if record.Kind == 2 {
		c.JSON(200, gin.H{
			"status": 0,
			"kind":   2,
			"msg":    service.GetFileDownloadLink(code, record.Filename),
		})
	}
}

func AddFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	res, err := service.SaveFile(file, c)
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

func AllRecords(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  0,
		"records": database.GetAllRecords(),
	})
}

func DeleteRecord(c *gin.Context) {
	code := c.Param("code")
	service.DeleteRecord(code)
	c.JSON(200, gin.H{
		"status": 0,
	})
}

func Usage(c *gin.Context) {
	c.String(http.StatusOK, `upload: curl -X POST http://THIS_SITE_ADDRESS/updown/file -F "file=@YOUR_FILE_PATH"
query:  curl http://THIS_SITE_ADDRESS/updown/record/YOUR_CODE`)
}
