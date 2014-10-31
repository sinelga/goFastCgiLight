package checkmetadata

import (
    "testing"
    "log/syslog"
    "fmt"
)
var golog, _ = syslog.New(syslog.LOG_ERR, "golog")


func TestGetmetadata(t *testing.T) {
	
	file := "/home/juno/test.html.gz"
	
//	var index bool = false
	
	ret :=Getmetadata(*golog, file)

	if ret == nil {
		
		fmt.Println("nil!!!!")
		
		
	}else {
		
		fmt.Println(ret[0],ret[1],ret[2])
		
	}

	

}

