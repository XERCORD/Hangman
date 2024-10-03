package affichage

func ReplaceWithUnderscores(mot string, lettreVisible rune) string {
	affichage := ""
	for _, lettre := range mot {
		if lettre == lettreVisible {
			affichage += string(lettre)
		} else if lettre == ' ' {
			affichage += " " // Conserver les espaces
		} else if lettre >= 'A' && lettre <= 'Z' || lettre >= 'a' && lettre <= 'z' {
			affichage += "_" // Remplacer uniquement les lettres alphabétiques par des underscores
		}
	}
	return affichage
}
