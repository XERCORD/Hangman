package affichage

import (
	contains "hangman/Travail_definitive/Contains"
)

func ReplaceWithMultipleLetters(mot string, indices []int) string {
	affichage := ""
	for i := 0; i < len(mot); i++ {
		if contains.Contains(indices, i) {
			affichage += string(mot[i]) // Révéler les lettres aux indices choisis
		} else if mot[i] == ' ' {
			affichage += " " // Conserver les espaces
		} else {
			affichage += "_" // Remplacer les autres lettres par des underscores
		}
	}
	return affichage
}
