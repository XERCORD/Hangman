package contains

func ContainsRune(mot string, lettre rune) bool {
	for _, l := range mot {
		if l == lettre {
			return true
		}
	}
	return false
}
