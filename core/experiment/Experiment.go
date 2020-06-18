package experiment

type Experiment struct {
	Id        int      `json:"expid"`
	Traffic   int      `json:"traffic"`
	Whitelist []string `json:"whitelists"`
	Tag       string   `json:"tag"`
	LogTag    string   `json:"logtag"`
	Buckets   Range    `json:"buckets"`
	WhiteSet  map[string]bool
}

type Range struct {
	start, end int
}

func (r Range) In(b int) bool {
	return r.start <= b && b < r.end
}

func (exp *Experiment) setBuckets(s, e int) {
	r := Range{
		start: s,
		end:   e,
	}
	exp.Buckets = r
}
