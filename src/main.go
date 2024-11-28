package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Fonction main
func main() {
	clearConsole()
	fmt.Println("Hello, welcome to the game.")
	if _, err := os.Stat("data/savePersonnage.json"); err == nil {
		// Le fichier existe, vérifie s'il est vide
		fileInfo, _ := os.Stat("data/savePersonnage.json")
		if fileInfo.Size() == 0 {
		} else {
			for {
				fmt.Println("A saved character exists. Do you want to load it? (y/n):")
				var choice string
				fmt.Scanln(&choice)

				if choice == "y" || choice == "Y" {
					// Charger le personnage à partir du fichier
					perso1, err := loadCharacterFromFile("data/savePersonnage.json")
					if err != nil {
						fmt.Println("Error loading character:", err)
						return
					}
					perso1.loading()
					return
				} else if choice == "n" || choice == "N" {
					err := ioutil.WriteFile("data/savePersonnage.json", []byte(""), 0644)
					if err != nil {
						fmt.Println("Error clearing saved character:", err)
					} else {
						fmt.Println("Saved character data cleared.")
						perso1 := CreatePersoMain()
						perso1.loading()
					}
					return
				} else {
					fmt.Println("Input error. Please enter 'y' or 'n'.")
				}
			}
		}
	}
	perso1 := CreatePersoMain()
	perso1.loading()
}

// Fonction pour créer le personnage
func CreatePersoMain() *Personnage {
	name, classe, pvMaxPerso := CreateChar()

	skillsMap, err := loadSkillsFromFile()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	inventaireInitial := []Inventory{
		{Elem: Item{Name: "Plume de Corbeau", Label: "plume_corbeau", Type: "items"}, Quantity: 1},
	}

	classSkills, ok := skillsMap[classe]
	if !ok {
		fmt.Println("Unknown class, no skills assigned.")
		return nil
	}

	skillInitial := make([]Skills, len(classSkills))
	copy(skillInitial, classSkills)

	perso1 := Init(
		name,
		classe,
		1,
		100,
		pvMaxPerso/2,
		pvMaxPerso,
		10,
		0,
		skillInitial,
		10,
		inventaireInitial,
	)
	return &perso1
}
