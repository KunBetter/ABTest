package context

import "encoding/json"

type ABContext struct {
	LayId      int               `json:"layId"`
	TraceTag   string            `json:"traceTag"`  //贯穿多层实验,记录实验Tag
	GlobalTag  string            `json:"globalTag"` //全局默认Tag
	ContextMap map[string]string `json:"contextMap"`
}

func (context *ABContext) toString() string {
	buf, _ := json.Marshal(context)
	return string(buf)
}
