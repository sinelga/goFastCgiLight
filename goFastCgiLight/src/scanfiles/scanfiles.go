package scanfiles

import (
	"log/syslog"
	"strings"
	
)

func Scan(golog syslog.Writer, path string) string {

	patharr := strings.Split(path, "/")

	site := patharr[9]
	
	
	return site
	
}

