package service

import (
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
	"updownloader-backend/config"
	"updownloader-backend/database"

	"github.com/flytam/filenamify"
	"github.com/gin-gonic/gin"
)

const chars string = "abcdefghijklmnopqrstuvwxyz0123456789"

func InitDir() {
	if err := os.MkdirAll(config.BasePath(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func randId() string {
	num := rand.Intn(36 * 36 * 36 * 36)
	p := make([]byte, 4)
	for i := 0; i < 4; i++ {
		p[i] = chars[num%36]
		num /= 36
	}
	return string(p)
}

func HasCode(code string) bool {
	// TODO: 这里改成数据库判定
	_, err := os.Stat(config.BasePath() + code)
	return err == nil
}

func SaveText(text string) (string, error) {
	var code string
	for i := 0; i < 10; i++ {
		ret := randId()
		if !HasCode(ret) {
			code = ret
			break
		}
	}

	if code == "" {
		return "", errors.New("cannot allocate code")
	}

	os.Mkdir(filepath.Join(config.BasePath(), code), os.ModePerm)
	fileName := code + ".txt"
	f, _ := os.Create(filepath.Join(config.BasePath(), code, fileName))
	defer f.Close()
	f.WriteString(text)
	database.InsertRecord(&database.Record{
		Code:       code,
		Kind:       1,
		Filename:   fileName,
		ExpireTime: time.Now().Add(time.Hour * 24),
	})
	return code, nil
}

func ReadText(code string) string {
	b, err := ioutil.ReadFile(filepath.Join(config.BasePath(), code, code+".txt"))
	if err != nil {
		return ""
	}
	return string(b)
}

func SaveFile(file *multipart.FileHeader, c *gin.Context) (string, error) {
	var code string
	for i := 0; i < 10; i++ {
		ret := randId()
		if !HasCode(ret) {
			code = ret
			break
		}
	}

	if code == "" {
		return "", errors.New("cannot allocate code")
	}

	os.Mkdir(filepath.Join(config.BasePath(), code), os.ModePerm)
	namifiedFilename, err := filenamify.Filenamify(file.Filename, filenamify.Options{})

	if err != nil {
		return "", err
	}

	err = c.SaveUploadedFile(file, filepath.Join(config.BasePath(), code, namifiedFilename))

	if err != nil {
		return "", err
	}

	database.InsertRecord(&database.Record{
		Code:       code,
		Kind:       2,
		Filename:   namifiedFilename,
		ExpireTime: time.Now().Add(time.Hour * 24),
	})

	return code, nil
}

func GetFileDownloadLink(code string, filename string) string {
	return "http://" + config.ListeningAddress() + "/webserver/" + code + "/" + filename
}

func DeleteRecord(code string) {
	database.DeleteRecordByCode(code)
	os.RemoveAll(filepath.Join(config.BasePath(), code))
}
