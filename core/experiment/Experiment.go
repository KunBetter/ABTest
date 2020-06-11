package experiment

type Experiment struct {
	Traffic   int            `json:"traffic"`
	Whitelist []string       `json:"whitelist"`
	Tag       string         `json:"tag"`
	LogTag    string         `json:"logtag"`
	Buckets   map[int]int    `json:"buckets"`
	WhiteSet  map[string]int `json:"whitelist"`
}

func (exp *Experiment) setBuckets(s, e int) {
	for i := s; i < e; i++ {
		exp.Buckets[i] = 1
	}
}
