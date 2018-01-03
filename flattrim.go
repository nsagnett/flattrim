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

func (f *flattrimizer) Flatten(value interface{}) map[string]interface{} {
	return f.FlattenWithPrefix("", value)
}

func (f *flattrimizer) FlattenWithPrefix(prefix string, value interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	f.flatten(result, prefix, value)
	return result
}

func (f *flattrimizer) flatten(result map[string]interface{}, prefix string, value interface{}) {

	if data, ok := value.(map[string]interface{}); ok {
		flattenKey := f.BuildFlattenKey(prefix)

		for key, val := range data {

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
	} else {
		if f.caseType == LOWERCASE {
			result[strings.ToLower(prefix)] = data
		} else {
			result[prefix] = data
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
