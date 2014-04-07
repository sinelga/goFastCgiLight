package prtitlegen

import (
	"comutils"
//	"math/rand"
//	"time"
//	"log"
)

func Generate(keywordsarr []string) []string {

	var titlearr []string
	for i := 0; i < 2; i++ {
	
		titlearr =append(titlearr,keywordsarr[comutils.Random(0, len(keywordsarr))])
	
	}
	
//	prtitle := keywordsarr[comutils.Random(0, len(keywordsarr))]
//	prtitle =comutils.UpcaseInitial(prtitle) +" " +keywordsarr[comutils.Random(0, len(keywordsarr))] +"."
	
	return titlearr

}
