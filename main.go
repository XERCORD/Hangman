package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	mots_faciles := []string{
		"chat",
		"chien",
		"table",
		"fleur",
		"livre",
		"arbre",
		"porte",
		"mer",
		"loup",
		"main",
	}

	mots_difficiles := []string{
		"voiture",
		"cahier",
		"maison",
		"jardin",
		"banane",
		"valise",
		"chapeau",
		"bouton",
		"dentiste",
		"clavier",
		"horloge",
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Print("==============================================================\n")
	fmt.Println("Vous souhaitez choisir un mot facile ou un mot difficile")
	fmt.Print("==============================================================\n")
	fmt.Println("1 - Un mot facile")
	fmt.Println("2 - Un mot difficile")
	fmt.Println("0 - Quitter")

	var choix int
	fmt.Print("Choix : ")
	fmt.Scan(&choix)

	var motAleatoire string
	var affichage string
	lettresProposees := make(map[rune]bool)
	switch choix {
	case 1:
		motAleatoire = mots_faciles[rand.Intn(len(mots_faciles))]
		lettreVisible := rune(motAleatoire[rand.Intn(len(motAleatoire))])
		affichage = replaceWithUnderscores(motAleatoire, lettreVisible)
		fmt.Println("Vous avez choisi un mot facile.")
		lettresProposees[lettreVisible] = true
	case 2:
		motAleatoire = mots_difficiles[rand.Intn(len(mots_difficiles))]
		indicesVisibles := rand.Perm(len(motAleatoire))[:2]
		affichage = replaceWithMultipleLetters(motAleatoire, indicesVisibles)
		fmt.Println("Vous avez choisi un mot diffiicile.")
		for _, i := range indicesVisibles {
			lettresProposees[rune(motAleatoire[i])] = true
		}
	case 0:
		fmt.Println("Vous avez choisi de quitter le jeu.")
		return
	default:
		fmt.Println("Choix invalide. Utilisation de la liste facile par défaut.")
		motAleatoire = mots_faciles[rand.Intn(len(mots_faciles))]
		lettreVisible := rune(motAleatoire[rand.Intn(len(motAleatoire))])
		affichage = replaceWithUnderscores(motAleatoire, lettreVisible)
		lettresProposees[lettreVisible] = true
	}

	vie := 9

	for {
		fmt.Print("================================\n")
		fmt.Printf("Le mot à deviner est : %s\n", affichage)
		fmt.Print("================================\n")
		// Demander à l'utilisateur d'entrer une lettre

		fmt.Println("Vous voulez proposer une lettre ou un mots entier ?")
		fmt.Println("1 - Proposer une lettre")
		fmt.Println("2 - Proposer un mot entier")

		var action int
		fmt.Print("Choix : ")
		fmt.Scan(&action)
		switch action {
		case 1:
			var lettreChoisie string
			fmt.Print("Entrez une lettre : ")
			fmt.Scan(&lettreChoisie) // Lire la lettre choisie
			fmt.Print("================================\n")

			lettreChoisie = strings.TrimSpace(lettreChoisie)
			if len(lettreChoisie) != 1 {
				fmt.Println("Veuillez entrer une seule lettre.")
				continue
			}

			// Prendre uniquement la première lettre entrée
			lettreChoisieRune := rune(lettreChoisie[0])

			// Vérifier si la lettre a déjà été proposée
			if lettresProposees[lettreChoisieRune] {
				fmt.Println("Cette lettre a déjà été proposée. Essayez une autre.")
				continue
			}
			// Ajouter la lettre à la liste des lettres proposées
			lettresProposees[lettreChoisieRune] = true

			// Mettre à jour l'affichage
			if containsRune(motAleatoire, lettreChoisieRune) {
				affichage = revealLetter(motAleatoire, affichage, lettreChoisieRune)
				fmt.Println("Bien joué !")
			} else {
				fmt.Println("Ce mot ne contient pas cette lettre.")
				vie--
				fmt.Printf("Il vous reste %d vies.\n", vie)
			}
		case 2:
			var motPropose string
			fmt.Print("Entrez un mot : ")
			fmt.Scan(&motPropose)

			motPropose = strings.TrimSpace(motPropose)

			// Vérifier si le mot proposé est correct
			if strings.EqualFold(motPropose, motAleatoire) {
				fmt.Printf("Félicitations, vous avez réussi à proposer le bon mot : %s\n", motAleatoire)
				fmt.Println("Appuyez sur Entrée pour terminer...")
				fmt.Scanln()
				break
			} else {
				fmt.Println("Ce n'est pas le bon mot.")
				vie -= 2
				fmt.Printf("Il vous reste %d vies.\n", vie)
			}
		default:
			fmt.Println("Choix invalide. Veuillez choisir 1 ou 2.")
			continue
		}

		// Vérifier si le mot est complètement deviné
		if !containsUnderscore(affichage) {
			fmt.Printf("Félicitations, vous avez deviné le mot : %s\n", motAleatoire)
			fmt.Println("Appuyez sur Entrée pour terminer...")
			fmt.Scanln()
			break
		}

		// Vérifier si les vies sont épuisées
		if vie <= 0 {
			fmt.Printf("Vous avez perdu ! Le mot était : %s\n", motAleatoire)
			fmt.Println("Appuyez sur Entrée pour terminer...")
			fmt.Scanln()
			break
		}
	}
}

func containsRune(mot string, lettre rune) bool {
	for _, l := range mot {
		if l == lettre {
			return true
		}
	}
	return false
}

// Autres fonctions (identiques à celles du code original)
func replaceWithUnderscores(mot string, lettreVisible rune) string {
	affichage := ""
	for _, lettre := range mot {
		if lettre == lettreVisible {
			affichage += string(lettre)
		} else {
			affichage += "_"
		}
	}
	return affichage
}

func replaceWithMultipleLetters(mot string, indices []int) string {
	affichage := ""
	for i := 0; i < len(mot); i++ {
		if contains(indices, i) {
			affichage += string(mot[i])
		} else {
			affichage += "_"
		}
	}
	return affichage
}

func contains(indices []int, val int) bool {
	for _, index := range indices {
		if index == val {
			return true
		}
	}
	return false
}

func revealLetter(mot string, affichage string, lettre rune) string {
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

func containsUnderscore(affichage string) bool {
	for _, l := range affichage {
		if l == '_' {
			return true
		}
	}
	return false
}
