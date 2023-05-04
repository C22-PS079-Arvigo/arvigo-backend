package utils

func InArray(needle interface{}, haystack []interface{}) bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false
}
