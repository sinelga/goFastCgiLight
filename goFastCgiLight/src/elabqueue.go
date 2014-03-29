package main

import (
	"domains"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"htmlfileexist"
	//	"io/ioutil"
	"log"
	"log/syslog"
	"math/rand"
	"time"
	//	"makenewsite"
	//	"os"
	//	"path"
)

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")

	}

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}

	//	golog.Info("elabque: Start firstpagebin")
	//	if qfirstpagesbin, err := redis.Strings(c.Do("HKEYS", "firstpagebin")); err != nil {
	//		golog.Crit(err.Error())
	//	} else {
	//
	//		for _, keystr := range qfirstpagesbin {
	//
	//			if webpagebytes, err := redis.Bytes(c.Do("HGET", "firstpagebin", keystr)); err != nil {
	//
	//				golog.Crit(err.Error())
	//
	//			} else {
	//
	//				dirpath := path.Dir(keystr)
	//				//				golog.Info("elabqueue:firstpagebin: dir " + dirpath)
	//
	//				if _, err := os.Stat(dirpath); err != nil {
	//					if os.IsNotExist(err) {
	//						// file does not exist
	//						//						golog.Info("elabqueue:firstpagebin: OK create dir " +dirpath)
	//						err := os.MkdirAll(dirpath, 0777)
	//						if err != nil {
	//							golog.Err("!!!elabqueue:firstpagebin: " + err.Error())
	//						}
	//
	//					} else {
	//						// other error
	//						golog.Err("!!!elabqueue:firstpagebin: " + err.Error())
	//					}
	//
	//				}
	//
	//				if err := ioutil.WriteFile(keystr, webpagebytes, 0666); err != nil {
	//					golog.Crit(err.Error())
	//				} else {
	//
	//					if _, err := redis.Int(c.Do("HDEL", "firstpagebin", keystr)); err != nil {
	//
	//						golog.Crit(err.Error())
	//
	//					}
	//
	//				}
	//
	//			}
	//		}
	//	}

	//	golog.Info("elabque: Start firstpage")
	//	if qfirstpages, err := redis.Int(c.Do("SCARD", "firstpage")); err != nil {
	//		golog.Crit(err.Error())
	//
	//	} else {
	//
	//		for i := 0; i < qfirstpages; i++ {
	//
	//			bfirstpage, _ := redis.Bytes(c.Do("SPOP", "firstpage"))
	//			var unmar domains.FirstPage
	//			err := json.Unmarshal(bfirstpage, &unmar)
	//			if err != nil {
	//
	//				golog.Crit(err.Error())
	//
	//			}
	//			makenewsite.Makenew(*golog, unmar)
	//
	//		}
	//
	//	}
	dir := "tiedotDB"
	rand.Seed(time.Now().UTC().UnixNano())

	tdDB, err := db.OpenDB(dir)
	defer tdDB.Close()
	if err != nil {
		panic(err)
	}

	col := tdDB.Use("Sites")

	golog.Info("elabque: Start pagetocreate")
	if qpages, err := redis.Int(c.Do("SCARD", "pagetocreate")); err != nil {
		log.Fatal(err)

	} else {

		for i := 0; i < qpages; i++ {

			msite, _ := redis.Bytes(c.Do("SPOP", "pagetocreate"))

			var unmar domains.SiteQue
			err := json.Unmarshal(msite, &unmar)
			if err != nil {
				golog.Crit(err.Error())
			}

			htmlfileexist.StartCheck(*golog, col, unmar.Locale, unmar.Themes, unmar.Domain, unmar.Pathinfo)

		}

	}

	golog.Info("elabque: END pagetocreate")
	c.Flush()
	c.Close()

}
