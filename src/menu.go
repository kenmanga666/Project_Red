package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
)

// Fonction pour clear la console
func clearConsole() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "darwin", "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("\n")
	}
}

// Fonction pour revenir en arrière dans le menu
func (p *Personnage) waitForEnter() {
	fmt.Println("Press Enter to go back...")
	fmt.Scanln()
	clearConsole()
}

// Fonction menu
func (p *Personnage) menu() {
	fmt.Println("1. View informations about your character")
	fmt.Println("2. View your inventory")
	fmt.Println("3. View your equipment")
	fmt.Println("4. View your skills")
	fmt.Println("5. Enter the shop")
	fmt.Println("6. Forgeron")
	fmt.Println("7. Combat")
	fmt.Println("8. Exit")
	fmt.Println("")

	var choix int
	fmt.Println("Enter your choice : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.menu()
	}

	switch choix {
	case 1:
		clearConsole()
		p.displayInfo()
		p.waitForEnter()
		p.menu()
	case 2:
		clearConsole()
		p.accessInventory()
		p.menu()
	case 3:
		clearConsole()
		p.accessEquipment()
		p.menu()
	case 4:
		clearConsole()
		p.checkSkills()
		p.waitForEnter()
		p.menu()
	case 5:
		clearConsole()
		p.enterShop()
		p.waitForEnter()
	case 6:
		clearConsole()
		p.forgeron()
	case 7:
		clearConsole()
		p.combatMenu()
	case 8:
		fmt.Println("Bye !")
		p.quitGame()
	default:
		fmt.Println("Invalid choice. Please choose a valid option.")
		p.waitForEnter()
		clearConsole()
		p.menu()
	}
}

// Fonction pour quitter le jeu
func (p *Personnage) quitGame() {
    err := p.saveCharacterToFile()
	if err != nil {
		fmt.Println("Error saving character data:", err)
	}

	os.Exit(0)
}