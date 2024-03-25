package easyget

import (
	"encoding/json"

	"github.com/oliveagle/jsonpath"
)

type JSONPathFilter struct {
	jp string
}

func NewJSONPathFilter(jp string) *JSONPathFilter {
	return &JSONPathFilter{jp: jp}
}

func (jpf *JSONPathFilter) Filt(d []byte) []byte {
	var jsonData interface{}
	err := json.Unmarshal(d, &jsonData)
	if err != nil {
		l.Errorf("Error unmarshalling JSON: %s", err)
	}
	res, err := jsonpath.JsonPathLookup(jsonData, jpf.jp)
	if err != nil {
		l.Errorf("Error applying JsonPath: %s", err)
		return []byte("")
	}

	filteredJSON, err := json.Marshal(res)
	if err != nil {
		l.Errorf("Error marshalling filtered JSON: %s", err)
		return []byte("")
	}

	return filteredJSON
}
