package utils

func StrInArray(arr []string, str string) bool {
	for _, x := range arr {
		if x == str {
			return true
		}
	}

	return false
}
