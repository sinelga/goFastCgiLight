package main

import (
	"domains"
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/garyburd/redigo/redis"
	"htmlfileexist"
	"log"
	"log/syslog"
	"math/rand"
	"time"

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
