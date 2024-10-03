package affichage

func RevealLetter(mot string, affichage string, lettre rune) string {
	newAffichage := ""
	for i, l := range mot {
		if l == lettre {
			newAffichage += string(l)
		} else {
			newAffichage += string(affichage[i])
		}
	}
	return newAffichage
}
