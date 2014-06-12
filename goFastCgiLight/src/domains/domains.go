package domains

import ()

type Site struct {
	Pathinfo   string
	Created    int64
	Updated    int64
	Hits       int
	Paragraphs []Paragraph
}

type Paragraph struct {
	Ptitle     string
	Pphrase    string
	Plocallink string
	Phost      string
	Sentences  []string
	Pushsite   string
}

type Htmlpage struct {
	Paragraphs []Paragraph
}

type SiteQue struct {
	Locale   string
	Themes   string
	Domain   string
	Pathinfo string
}

type Domaincsv struct {
	Locale string
	Themes string
	Domain string
	Ip     string
}
