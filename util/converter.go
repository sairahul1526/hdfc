package util

// ExtractValuesFromArrayMap -
func ExtractValuesFromArrayMap(data []map[string]interface{}, key string) []string {
	keys := []string{}
	for _, object := range data {
		keys = append(keys, object[key].(string))
	}
	return keys
}
