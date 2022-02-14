package crontab

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
	"updownloader-backend/config"
	"updownloader-backend/database"
	"updownloader-backend/service"

	"github.com/robfig/cron"
)

func timedDelete() {
	log.Println("timed delete")
	rd, err := ioutil.ReadDir(config.BasePath())
	if err != nil {
		log.Println(err)
		return
	}
	for _, fi := range rd {
		if fi.IsDir() {
			code := fi.Name()
			if record, exi := database.QueryRecordByCode(code); !exi {
				os.RemoveAll(filepath.Join(config.BasePath(), code))
				fmt.Println("removed", code, "because it does not exist in database")
			} else {
				if record.ExpireTime.Before(time.Now()) {
					service.DeleteRecord(code)
					fmt.Println("removed", code, "because it expired")
				}
			}
		}
	}
}

func InitCron() {
	timedDelete()
	c := cron.New()
	c.AddFunc("@daily", timedDelete)
	c.Start()
}
