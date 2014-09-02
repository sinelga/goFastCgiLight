package startones

import (
		"log"
	"log/syslog"
	"testing"
	

)

func TestStart(t *testing.T) {
	
	golog, err := syslog.New(syslog.LOG_ERR, "golog")

	defer golog.Close()
	if err != nil {
		log.Fatal("error writing syslog!!")
	}
	parameters :=Start(*golog)
	
	golog.Info(parameters[1])
	
	
	
}