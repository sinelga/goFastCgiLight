package pipelingpush

import (
	"domains"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log/syslog"
	"os"
)

func PushInQueue(golog syslog.Writer, queuesys string, pushtoQueueArr [][]string) {

	var sitejsonarr [][]byte

	for _, pushstr := range pushtoQueueArr {

		site := domains.SiteQue{

			Locale:   pushstr[0],
			Themes:   pushstr[1],
			Domain:   pushstr[2],
			Pathinfo: pushstr[3],
		}
		bsite, _ := json.Marshal(site)

		sitejsonarr = append(sitejsonarr, bsite)

		c, err := redis.Dial("tcp", ":6379")
		defer c.Close()
		if err != nil {

			golog.Err("PushInQueue: " + err.Error())
			os.Exit(1)
		} else {

			_,err := c.Do("MULTI")
			if err != nil {
				golog.Err("PushInQueue: " + err.Error())
			}

			for _, hit := range sitejsonarr {

				c.Send("SADD", "pagetocreate", hit)

			}
			_, err = c.Do("EXEC")
			if err != nil {
				golog.Err("PushInQueue: " + err.Error())
			}
			
			c.Flush()
			c.Close()

		}

	}

}
