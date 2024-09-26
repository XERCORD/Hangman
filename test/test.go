package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	motsFrancais1 := []string{
		"chat",
		"chien",
		"livre",
		"fleur",
		"porte",
		"arbre",
		"soleil",
		"rêve",
		"lune",
		"table",
	}
	motsFrancais2 := []string{
		"chanson",
		"bananes",
		"forêtier",
		"cascade",
		"cerises",
		"giraffe",
		"ballons",
		"plantes",
		"horizon",
		"musique",
	}

	rand.Seed(time.Now().UnixNano())

	// Demander à l'utilisateur de choisir une difficulté
	var choix int
	fmt.Println("Entrez une difficulté :")
	fmt.Println("(1) Facile")
	fmt.Println("(2) Difficile")
	fmt.Scan(&choix)

	var motAleatoire string
	var affichage string

	// Sélectionner la liste de mots selon le choix de l'utilisateur
	switch choix {
	case 1:
		motAleatoire = motsFrancais1[rand.Intn(len(motsFrancais1))]
		// Afficher une lettre aléatoire
		lettreVisible := rune(motAleatoire[rand.Intn(len(motAleatoire))])
		affichage = replaceWithUnderscores(motAleatoire, lettreVisible)
	case 2:
		motAleatoire = motsFrancais2[rand.Intn(len(motsFrancais2))]
		// Afficher deux lettres aléatoires
		indicesVisibles := rand.Perm(len(motAleatoire))[:2]
		affichage = replaceWithMultipleLetters(motAleatoire, indicesVisibles)
	default:
		fmt.Println("Choix invalide. Utilisation de la liste facile par défaut.")
		motAleatoire = motsFrancais1[rand.Intn(len(motsFrancais1))]
		lettreVisible := rune(motAleatoire[rand.Intn(len(motAleatoire))])
		affichage = replaceWithUnderscores(motAleatoire, lettreVisible)
	}
	vie := 9
	// Boucle pour deviner les lettres
	for {
		fmt.Printf("Le mot à deviner est : %s\n", affichage)

		// Demander à l'utilisateur d'entrer une lettre
		var lettreDevinee rune
		fmt.Print("Entrez une lettre : ")
		fmt.Scanf("%c\n", &lettreDevinee) // Lire une seule lettre

		// Mettre à jour l'affichage
		if containsRune(motAleatoire, lettreDevinee) {
			affichage = revealLetter(motAleatoire, affichage, lettreDevinee)
			fmt.Println("Bien joué !")
		} else {
			fmt.Println("Ce mot ne contient pas cette lettre")
			vie--
			fmt.Println("il vous reste %d vie", vie)
		}

		// Vérifier si le mot est complètement deviné
		if !containsUnderscore(affichage) {
			fmt.Printf("Félicitations, vous avez deviné le mot : %s\n", motAleatoire)
			break
		}
		if vie <= 0 {
			fmt.Printf("Vous ne pouvez plus entrer de lettres car vos vie sont épuisé : %s\n", motAleatoire)
			break
		}
	}

	// Ajouter une pause pour garder la fenêtre ouverte
	fmt.Println("Appuyez sur Entrée pour terminer...")
	fmt.Scanln() // Attendre une seconde entrée pour garder le programme ouvert
}

// Fonction pour remplacer les lettres restantes par des underscores
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

// Fonction pour afficher deux lettres aléatoires
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

// Fonction pour vérifier si un indice est dans une liste
func contains(indices []int, val int) bool {
	for _, index := range indices {
		if index == val {
			return true
		}
	}
	return false
}

// Fonction pour vérifier si une rune est dans un mot
func containsRune(mot string, lettre rune) bool {
	for _, l := range mot {
		if l == lettre {
			return true
		}
	}
	return false
}

// Fonction pour dévoiler une lettre dans l'affichage
func revealLetter(mot string, affichage string, lettre rune) string {
	newAffichage := ""
	for i, l := range mot {
		if l == lettre {
			newAffichage += string(l)
		} else {
			newAffichage += string(affichage[i]) // garder le même caractère
		}
	}
	return newAffichage
}

// Fonction pour vérifier s'il reste des underscores
func containsUnderscore(affichage string) bool {
	for _, l := range affichage {
		if l == '_' {
			return true
		}
	}
	return false
}
