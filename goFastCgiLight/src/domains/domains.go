package domains

import (
	"time"
)

type Site struct {


	Pathinfo string
	Created  int
	Updated  int
	Hits     int
}

type Paragraph struct {
	Created    time.Time
	Ptitle     string
	Pphrase    string
	Plocallink string
	Phost      string
	Sentences  []string
}
