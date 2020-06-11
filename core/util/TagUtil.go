package util

import (
	"strconv"
	"strings"
)

func AppendTag(traceTag string, layId int, tag string) string {
	if len(traceTag) == 0 {
		return strconv.Itoa(layId) + "&" + tag
	}

	return traceTag + "," + strconv.Itoa(layId) + "&" + tag
}

func SplitTraceTag(traceTag string) map[string]string {
	if len(traceTag) == 0 {
		return nil
	}

	tagMap := make(map[string]string)

	tags := strings.Split(traceTag, ",")
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		kv := strings.Split(tag, "&")
		if 2 == len(kv) {
			tagMap[kv[0]] = kv[1]
		}
	}

	return tagMap
}
