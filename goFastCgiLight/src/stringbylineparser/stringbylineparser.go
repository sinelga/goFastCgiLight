package stringbylineparser

import (
//	"fmt"
	"strings"
)

func Parse(sentenses string) []string {

	//	fmt.Println(sentenses)
	var ret_arrstrings []string

	lines := strings.Split(sentenses, "\n")

	for _, line := range lines {

		if len(line) > 6 {
//			fmt.Println(strings.TrimSpace(line))
			
			ret_arrstrings = append(ret_arrstrings,strings.TrimSpace(line)) 
		}

	}

	return ret_arrstrings

}
