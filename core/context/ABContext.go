package context

import "encoding/json"

type ABContext struct {
	LayId      int               `json:"layId"`
	TraceTag   string            `json:"traceTag"` //Throughout the multi-layer experiment
	GlobalTag  string            `json:"globalTag"`
	ContextMap map[string]string `json:"contextMap"`
}

func (context *ABContext) toString() string {
	buf, _ := json.Marshal(context)
	return string(buf)
}
