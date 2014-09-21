package scanfiles

import (
	"log/syslog"
//	"github.com/garyburd/redigo/redis"
	"strings"
//	"fmt"
	

)

func Scan(golog syslog.Writer, path string) string {

	patharr := strings.Split(path, "/")

	site := patharr[9]
	
	
	return site
	
}

