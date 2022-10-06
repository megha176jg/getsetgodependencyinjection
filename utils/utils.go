package utils

func ContainString(l []string, s string) bool {
	for _, i := range l {
		if i == s {
			return true
		}
	}
	return false
}
