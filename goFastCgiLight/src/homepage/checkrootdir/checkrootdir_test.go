package checkrootdir

import (
    "testing"
    "log/syslog"
    "fmt"
)


var golog, _ = syslog.New(syslog.LOG_ERR, "golog")

func TestCheck(t *testing.T) {
	
	var testArr =[]string{"/home/juno/git/goFastCgiLight/goFastCgiLight/www/it_IT/finance/test.com/index.html",
		"/home/juno/git/goFastCgiLight/goFastCgiLight/www/it_IT/finance/test.com/test.html.gz",
		"/home/juno/git/goFastCgiLight/goFastCgiLight/www/it_IT/finance/test.com/tesss12345678/test2/index.html",
		"/home/juno/git/goFastCgiLight/goFastCgiLight/www/it_IT/finance/test.com/tesss12345678/test2/test.html.gz",
				
	}
	
	for _,testfile :=range testArr {
		
		fmt.Println(testfile)
		siteinfo := Check(*golog,testfile )
		
		for _,site :=range siteinfo {
			
			fmt.Println(site)
			
		} 
		
		
	}

}

