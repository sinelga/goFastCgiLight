package main

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"htmlfileexist"
	"log"
	"log/syslog"
	"startones"

)

var startparameters []string

func main() {

	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}

	startparameters=startones.Start(*golog)

	c, err := redis.Dial("tcp", ":6379")
	if err != nil {

		golog.Crit(err.Error())
	}

	if qpages, err := redis.Int(c.Do("SCARD", "pagetocreate")); err != nil {
	
		golog.Crit(err.Error())

	} else {

		golog.Info("elabque: Start pagetocreate elaborate "+string(qpages) )
		for i := 0; i < qpages; i++ {

			msite, _ := redis.Bytes(c.Do("SPOP", "pagetocreate"))
			c.Flush()

			var unmar domains.SiteQue
			err := json.Unmarshal(msite, &unmar)
			if err != nil {
				golog.Crit(err.Error())
			}

			htmlfileexist.StartCheckNoDB(*golog, unmar.Locale, unmar.Themes, unmar.Domain, unmar.Pathinfo,startparameters)

		}

	}
	
	golog.Info("elabque: END pagetocreate")
	c.Flush()
	c.Close()

}
