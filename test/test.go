package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var pendu = []string{
	`
      
      
      
      
      
     ===`, // État initial vide
	`
  +---+
      |
      |
      |
      |
     ===`,
	`
  +---+
  O   |
      |
      |
      |
     ===`,
	`
  +---+
  O   |
  |   |
      |
      |
     ===`,
	`
  +---+
  O   |
 /|   |
      |
      |
     ===`,
	`
  +---+
  O   |
 /|\  |
      |
     ===`,
	`
  +---+
  O   |
 /|\  |
 /    |
     ===`,
	`
  +---+
  O   |
 /|\  |
 / \  |
     ===`,
	`
  +---+
 [O   |
 /|\  |
 / \  |
     ===`,
	`
  +---+
 [O]  |
 /|\  |
 / \  |
     ===`,
}

func main() {
	// Lire les fichiers et afficher les mots depuis des fichiers txt en utilisant des byte
	mots_faciles, err := lireFichierAvecBytes("mots_faciles.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier 'mots_faciles.txt':", err)
		return
	}

	mots_difficiles, err := lireFichierAvecBytes("mots_difficiles.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier 'mots_difficiles.txt':", err)
		return
	}

	// Interface utilisateur
	rand.Seed(time.Now().UnixNano())
	fmt.Print("==============================================================\n")
	fmt.Println("Vous souhaitez choisir un mot facile ou un mot difficile")
	fmt.Print("==============================================================\n")
	fmt.Println("1 - Un mot facile")
	fmt.Println("2 - Un mot difficile")
	fmt.Println("0 - Quitter")

	var choix int
	for {
		fmt.Print("Choix : ")
		fmt.Scan(&choix)
		if choix == 1 || choix == 2 || choix == 0 {
			break
		} else {
			fmt.Println("Choix invalide. Veuillez entrer 1, 2, ou 0.")
		}
	}

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
		fmt.Println("Vous avez choisi un mot difficile.")
		for _, i := range indicesVisibles {
			lettresProposees[rune(motAleatoire[i])] = true
		}
	case 0:
		fmt.Println("Vous avez choisi de quitter le jeu.")
		return
	}

	vie := 9
	for {
		fmt.Print("================================\n")
		fmt.Printf("Le mot à deviner est : %s\n", affichage)
		fmt.Print("================================\n")
		fmt.Println(pendu[9-vie]) // Affiche le dessin du pendu en fonction du nombre de vies restantes
		fmt.Println("Vous voulez proposer une lettre ou un mot entier ?")
		fmt.Println("1 - Proposer une lettre")
		fmt.Println("2 - Proposer un mot entier")
		fmt.Println("0 - Quitter")

		var action int
		fmt.Print("Choix : ")
		fmt.Scan(&action)

		if action != 1 && action != 2 && action != 0 {
			fmt.Println("Choix invalide. Veuillez entrer 1, 2 ou 0.")
			continue
		}

		switch action {
		case 1:
			var lettreChoisie string
			fmt.Print("Entrez une lettre : ")
			fmt.Scan(&lettreChoisie)
			fmt.Print("================================\n")

			lettreChoisie = strings.TrimSpace(lettreChoisie)
			if len(lettreChoisie) != 1 {
				fmt.Println("Veuillez entrer une seule lettre.")
				continue
			}

			lettreChoisieRune := rune(lettreChoisie[0])

			if lettresProposees[lettreChoisieRune] {
				fmt.Println("Cette lettre a déjà été proposée. Essayez une autre.")
				continue
			}

			lettresProposees[lettreChoisieRune] = true

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

			if strings.EqualFold(motPropose, motAleatoire) {
				fmt.Printf("Félicitations, vous avez proposé le bon mot : %s\n", motAleatoire)
				fmt.Println("Appuyez sur Entrée pour terminer...")
				fmt.Scanln()
				return
			} else {
				fmt.Println("Ce n'est pas le bon mot.")
				vie -= 2
				fmt.Printf("Il vous reste %d vies.\n", vie)
			}
		case 0:
			return
		}

		if !containsUnderscore(affichage) {
			fmt.Printf("Félicitations, vous avez deviné le mot : %s\n", motAleatoire)
			fmt.Scanln()
			break
		}

		if vie <= 0 {
			fmt.Println(pendu[9]) // Affiche le pendu complet lors de la défaite
			fmt.Printf("Vous avez perdu ! Le mot était : %s\n", motAleatoire)
			fmt.Scanln()
			break
		}
	}
}

func lireFichierAvecBytes(nomFichier string) ([]string, error) {
	fileData, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}

	var mots []string
	word := []byte{}
	breakLine := []byte("\n")

	for _, data := range fileData {
		if !bytes.Equal([]byte{data}, breakLine) {
			word = append(word, data)
		} else {
			mots = append(mots, string(word))
			word = word[:0]
		}
	}

	if len(word) > 0 {
		mots = append(mots, string(word))
	}

	return mots, nil
}

func containsRune(mot string, lettre rune) bool {
	for _, l := range mot {
		if l == lettre {
			return true
		}
	}
	return false
}

func replaceWithUnderscores(mot string, lettreVisible rune) string {
	affichage := ""
	for _, lettre := range mot {
		if lettre == lettreVisible {
			affichage += string(lettre)
		} else if lettre == ' ' {
			affichage += " " // Conserver les espaces
		} else {
			affichage += "_" // Remplacer uniquement les caractères visibles
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
