package contains

func ContainsUnderscore(affichage string) bool {
	for _, l := range affichage {
		if l == '_' {
			return true
		}
	}
	return false
}
