package utils

// ConvertArrayToMap converts an interface array to Map with key as interface and value as its index
func ConvertArrayToMap(inArray []interface{}) map[interface{}]int {
	arrMap := make(map[interface{}]int, len(inArray))
	for index, element := range inArray {
		arrMap[element] = index
	}
	return arrMap
}
