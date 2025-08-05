package util

func IntInList(i int, list []int) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}
