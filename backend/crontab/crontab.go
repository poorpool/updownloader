package crontab

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
	"updownloader-backend/database"

	"github.com/robfig/cron"
)

func timedDelete() {
	log.Println("try delete")
	rd, err := ioutil.ReadDir("/data/updownloader")
	if err != nil {
		log.Println(err)
		return
	}
	for _, fi := range rd {
		if fi.IsDir() {
			code := fi.Name()
			if record, exi := database.QueryRecordByCode(code); !exi {
				os.RemoveAll(filepath.Join("/data/updownloader", code))
				fmt.Println("removed", code, "because it does not exist in database")
			} else {
				if record.ExpireTime.Before(time.Now()) {
					os.RemoveAll(filepath.Join("/data/updownloader", code))
					database.DeleteRecordByCode(code)
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
