package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fonction menu combat
func (p *Personnage) combatMenu() {
	fmt.Println("What do you want to do ?")
	fmt.Println("Be careful you won't be able to go back until the end of the fight once you've chosen an option.")
	fmt.Println("1. Training fight")
	fmt.Println("2. Mission fight")
	fmt.Println("0. Back")
	fmt.Println("")
	var choix int
	fmt.Println("Enter your choice :")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.menu()
	}
	switch choix {
	case 1:
		clearConsole()
		p.trainingFight()
	case 2:
		clearConsole()
		p.missionFight()
	case 0:
		clearConsole()
		p.menu()
	default:
		fmt.Println("Input error.")
		p.combatMenu()
	}
}

// Fonction pour l'initiative
func Initiative() int {
	return rand.Intn(20) + 1
}

// Fonction pour le training
func (p *Personnage) trainingFight() {
	PinitialPV := p.pv
	var m monster
	m.InitGoblin()
	Minitiative := Initiative()
	Pinitiative := Initiative()
	if Pinitiative >= Minitiative {
		fmt.Println("You have an initiative advantage.")
		fmt.Println("You attack first.")
		fmt.Println("")
	} else {
		fmt.Println(m.Name, "has an initiative advantage.")
		fmt.Println(m.Name, "attacks first.")
		fmt.Println("")
	}
	for tour := 1; m.PV > 0 && p.checkPV() > 0; tour++ {
		if Pinitiative > Minitiative {
			fmt.Println("Turn", tour)
			fmt.Println("Your turn")
			p.attack(&m)
			fmt.Println(m.Name, "remaining life points", m.PV, "/", m.PVMax)
			fmt.Println("")
			if m.PV <= 0 {
				break
			}
			fmt.Println(m.Name, "turn")
			m.GoblinPattern(tour, p)
			fmt.Println("")
		} else {
			fmt.Println("Turn", tour)
			fmt.Println(m.Name, "turn")
			m.GoblinPattern(tour, p)
			fmt.Println("")
			if p.checkPV() <= 0 {
				break
			}
			fmt.Println("Your turn")
			p.attack(&m)
			fmt.Println(m.Name, "remaining life points", m.PV, "/", m.PVMax)
			fmt.Println("")
		}
	}
	if p.dead() {
		fmt.Println("You Died ! Back to the main menu.")
		p.waitForEnter()
		p.menu()
	} else {
		fmt.Println("You won ! Back to the main menu.")
		p.pv = PinitialPV
		p.waitForEnter()
		p.menu()
	}
}

// Fonction pour le combat
func (p *Personnage) missionFight() {
	var choix int
	var choosenmonster string
	fmt.Println("Which monster do you want to fight ?")
	fmt.Println("1. Goblin")
	fmt.Println("2. Orc")
	fmt.Println("3. ????")
	fmt.Println("Enter your choice :")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.menu()
	}
	switch choix {
	case 1:
		choosenmonster = "Goblin"
	case 2:
		choosenmonster = "Orc"
	case 3:
		choosenmonster = "SecretBoss"
	default:
		fmt.Println("Input error.")
		p.missionFight()
	}
	var m monster
	m.InitMonster(choosenmonster)
	clearConsole()
	if choosenmonster == "SecretBoss" {
		SecretBossEntrance()
	}
	fmt.Println("You are fighting", m.Name)
	fmt.Println("Player's PV", p.pv, "/", p.pvMax)
	fmt.Println(m.Name, "PV", m.PV, "/", m.PVMax)
	Minitiative := Initiative()
	Pinitiative := Initiative()
	if Pinitiative >= Minitiative {
		fmt.Println("You have an initiative advantage.")
		fmt.Println("You attack first.")
		fmt.Println("")
	} else {
		fmt.Println(m.Name, "has an initiative advantage.")
		fmt.Println(m.Name, "attacks first.")
		fmt.Println("")
	}
	for tour := 1; m.PV > 0 && p.pv > 0; tour++ {
		if Pinitiative > Minitiative {
			fmt.Println("Turn", tour)
			fmt.Println("Your turn")
			p.attackMenu(&m)
			fmt.Println(m.Name, "remaining life points", m.PV, "/", m.PVMax)
			fmt.Println("")
			if m.PV <= 0 {
				break
			}
			fmt.Println(m.Name, "turn")
			m.MonsterPattern(tour, p)
			fmt.Println("")
		} else {
			fmt.Println("Turn", tour)
			fmt.Println(m.Name, "turn")
			m.MonsterPattern(tour, p)
			fmt.Println("")
			if p.pv <= 0 {
				break
			}
			fmt.Println("Your turn")
			p.attackMenu(&m)
			fmt.Println(m.Name, "remaining life points", m.PV, "/", m.PVMax)
			fmt.Println("")
		}
	}
	if p.dead() {
		fmt.Println("You Died ! Back to the main menu.")
		p.waitForEnter()
		clearConsole()
		p.menu()
	} else {
		fmt.Println("You won !")
		p.reward(&m)
		fmt.Println("Back to the main menu.")
		p.waitForEnter()
		clearConsole()
		p.menu()
	}
}

// Fonction pour le menu d'attaque
func (p *Personnage) attackMenu(m *monster) {
	fmt.Println("What do you want to do ?")
	fmt.Println("1. Attack")
	fmt.Println("2. Use skill")
	fmt.Println("3. Use Item")
	fmt.Println("")
	var choix int
	fmt.Println("Enter your choice : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error. Please use a number.")
		p.menu()
	}
	switch choix {
	case 1:
		clearConsole()
		p.attack(m)
	case 2:
		clearConsole()
		p.specialAttack(m)
	case 3:
		clearConsole()
		used := p.accessFightInventory(m)
		if !used {
			p.attackMenu(m)
		}
	default:
		fmt.Println("Input error.")
		p.attackMenu(m)
	}
}

// Fonction pour l'attaque
func (p *Personnage) attack(m *monster) {
	fmt.Println(p.nom, "deals to", m.Name, p.Attack, "damages")
	m.PV -= p.Attack
}

// Fonction pour l'attaque avec les skills
func (p *Personnage) specialAttack(m *monster) {
	fmt.Println("Which skill do you want to use ?")
	for i := 0; i < len(p.skills); i++ {
		ManaCost := p.skills[i].ManaCost
		for j := 1; j < p.skills[i].Level; j++ {
			ManaCost = ManaCost + 6
		}
		fmt.Printf("%d. %s Level %v required mana: %v\n", i+1, p.skills[i].Name, p.skills[i].Level, ManaCost)
	}
	fmt.Println("0. Back")
	var choix int
	fmt.Print("Enter the skill number you want to use : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		return
	}
	switch {
	case choix == 0:
		fmt.Println("Canceled the special attack.")
		p.attackMenu(m)
	case choix >= 1 && choix <= len(p.skills):
		clearConsole()
		skillIndex := choix - 1
		selectedSkill := p.skills[skillIndex]
		p.UseSkill(selectedSkill, m)
	default:
		fmt.Println("Invalid choice. Please choose a valid skill number.")
	}
}

// Fonction pour les récompenses
func (p *Personnage) reward(m *monster) {
	if m.Name == "Goblin" {
		fmt.Println("You won", 20, "XP")
		p.addXP(20)
		p.LevelUp()
		fmt.Println("You won", 10, "Gold")
		p.addGold(10)
	} else if m.Name == "Orc" {
		fmt.Println("You won", 30, "XP")
		p.addXP(30)
		p.LevelUp()
		fmt.Println("You won", 15, "Gold")
		p.addGold(15)
	} else if m.Name == "World Eater" {
		fmt.Println("You won", 5000, "XP")
		p.addXP(5000)
		p.LevelUp()
		fmt.Println("You won", 1000, "Gold")
		p.addGold(1000)
	}
}

// Fonction pour l'entrée du Secret Boss
func SecretBossEntrance() {
	fmt.Println("You've found the Secret Boss !")
	fmt.Println("He have random Critical Chance and Critical Damages each turn, Be careful !")
	fmt.Println("Warning when Boss's Critical Chance is 10, it's a critical hit !")
	fmt.Println("Warning Boss's Critical Damages is a damage multiplier for when he do a Critical Hit !")
	fmt.Println("Warning when Boss's deal damages to you, he restore 20%")
	fmt.Println("of his attack in life points ! And 30% if it's a critical hit !")
	fmt.Println("Good luck adventurer !")
	time.Sleep(10 * time.Second)
}
