package contains

func Contains(indices []int, val int) bool {
	for _, index := range indices {
		if index == val {
			return true
		}
	}
	return false
}
