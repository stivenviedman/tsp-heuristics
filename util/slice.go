package util

// returns true if slice of strings contains target string
func Contains(s []string, str string) bool {
	for _, value := range s {
		if str == value {
			return true
		}
	}

	return false
}
