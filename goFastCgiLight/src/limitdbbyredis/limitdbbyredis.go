package limitdbbyredis

import (
	"github.com/garyburd/redigo/redis"
	"log/syslog"
)

func MakeNewDb(golog syslog.Writer, c redis.Conn, sitesmap map[string]int) {

	for site, pages := range sitesmap {

		golog.Info(site)

		if _, err := c.Do("ZADD", "limitsites", pages, site); err != nil {
			golog.Err("limitdbbyredis: " + err.Error())

		}

	}

}
