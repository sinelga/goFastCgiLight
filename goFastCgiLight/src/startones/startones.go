package startones

import (
	"io/ioutil"
	"log/syslog"
	"strings"
)

func Start(golog syslog.Writer) []string {


	content, err := ioutil.ReadFile("/home/juno/git/goFastCgiLight/goFastCgiLight/config.txt")
	if err != nil {
		//Do something
		golog.Err(err.Error())
	}
	parameters := strings.Split(string(content), ",")
	cleanparameters :=[]string{strings.TrimSpace(parameters[0]),strings.TrimSpace(parameters[1])}
	
	golog.Info("startones redis "+parameters[0]+" "+parameters[1])
	

	return cleanparameters 

}
