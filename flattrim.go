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
	binder   string
}

func NewFlattrimizer(caseType CaseType) *flattrimizer {
	return &flattrimizer{caseType: caseType, binder: "."}
}

func (f *flattrimizer) SetBinder(binder string) {
	f.binder = binder
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

	flattenKey := f.buildFlattenKey(prefix)

	for key, val := range value {

		switch newVal := val.(type) {
		case map[string]interface{}:
			f.flatten(result, fmt.Sprint(flattenKey, key), newVal)
		case []interface{}:
			arrayObjects := make([]interface{}, 0)
			for _, obj := range newVal {
				if v, ok := obj.(map[string]interface{}); ok {
					arrayObj := make(map[string]interface{})
					f.flatten(arrayObj, "", v)
					arrayObjects = append(arrayObjects, arrayObj)
				} else {
					arrayObjects = append(arrayObjects, obj)
				}
			}
			if f.caseType == LOWERCASE {
				result[strings.ToLower(fmt.Sprint(flattenKey, key))] = arrayObjects
			} else {
				result[fmt.Sprint(flattenKey, key)] = arrayObjects
			}
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

func (f *flattrimizer) buildFlattenKey(prefix string) string {
	if prefix == "" {
		return prefix
	} else {
		return fmt.Sprint(prefix, f.binder)
	}
}
