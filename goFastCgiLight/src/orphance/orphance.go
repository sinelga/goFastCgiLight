package orphance

import (
	"log/syslog"
	"os"
)

func LookUp(golog syslog.Writer, path string) {
	
	if err := os.Remove(path); err != nil {
		
		golog.Err(err.Error())
	
	}
	
}
