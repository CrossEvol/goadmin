package utils

import "encoding/json"

func ToJson(v any) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "{}"
	}
	return string(bytes)
}

func ToPrettyJson(v any) string {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(bytes)
}
