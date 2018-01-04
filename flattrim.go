package flattrim

import (
	"fmt"
	"strings"
)

type CaseType int

const (
	LOWERCASE CaseType = 1 + iota
	KEEPCASE
)

type flattrimizer struct {
	caseType CaseType
}

func NewFlattrimizer(caseType CaseType) *flattrimizer {
	return &flattrimizer{caseType: caseType}
}

func (f *flattrimizer) SetCaseType(caseType CaseType) {
	f.caseType = caseType
}

func (f *flattrimizer) Flatten(value map[string]interface{}) map[string]interface{} {
	return f.FlattenWithPrefix("", value)
}

func (f *flattrimizer) FlattenWithPrefix(prefix string, value map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	f.flatten(result, prefix, value)
	return result
}

func (f *flattrimizer) flatten(result map[string]interface{}, prefix string, value map[string]interface{}) {

	flattenKey := f.BuildFlattenKey(prefix)

	for key, val := range value {

		switch newVal := val.(type) {
		case map[string]interface{}:
			f.flatten(result, fmt.Sprint(flattenKey, key), newVal)
		case []map[string]interface{}:
			arrayObjects := make([]map[string]interface{}, 0)

			for _, obj := range newVal {
				arrayObj := make(map[string]interface{})
				f.flatten(arrayObj, "", obj)
				arrayObjects = append(arrayObjects, arrayObj)
			}
			if f.caseType == LOWERCASE {
				result[strings.ToLower(fmt.Sprint(flattenKey, key))] = arrayObjects
			} else {
				result[fmt.Sprint(flattenKey, key)] = arrayObjects
			}
		default:
			if f.caseType == LOWERCASE {
				result[strings.ToLower(fmt.Sprint(flattenKey, key))] = val
			} else {
				result[fmt.Sprint(flattenKey, key)] = val
			}
		}
	}
}

func (f *flattrimizer) BuildFlattenKey(prefix string) string {
	if prefix == "" {
		return prefix
	} else {
		return fmt.Sprint(prefix, ".")
	}
}
