package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Structure de l'inventaire
type Inventory struct {
	Elem     Item
	Quantity int
}

// Fonction pour équiper une armure
func (p *Personnage) equipArmor(equipement Item) {
	if equipement.SubType == "Head" {
		if p.equipements.head.Name != "" {
			p.unequipArmor(p.equipements.head)
		}
		p.equipements.head = equipement
		p.pvMax += equipement.BonusLife
	} else if equipement.SubType == "Body" {
		if p.equipements.body.Name != "" {
			p.unequipArmor(p.equipements.body)
		}
		p.equipements.body = equipement
		p.pvMax += equipement.BonusLife
	} else if equipement.SubType == "Boots" {
		if p.equipements.boots.Name != "" {
			p.unequipArmor(p.equipements.boots)
		}
		p.equipements.boots = equipement
		p.pvMax += equipement.BonusLife
	}
	p.removeInventoryItem(equipement.Name, 1)
	fmt.Println("You have equipped ", equipement.Name, ", +", equipement.BonusLife, "PVMax")
}

// Fonction pour déséquiper une armure
func (p *Personnage) unequipArmor(equipement Item) {
	if len(p.inventaire) >= p.inventoryCapacity {
		fmt.Println("The inventory is full. You cannot unequip this object.")
		return
	}
	if equipement.SubType == "Head" {
		p.equipements.head = Item{}
		p.pvMax -= equipement.BonusLife
	} else if equipement.SubType == "Body" {
		p.equipements.body = Item{}
		p.pvMax -= equipement.BonusLife
	} else if equipement.SubType == "Boots" {
		p.equipements.boots = Item{}
		p.pvMax -= equipement.BonusLife
	}
	p.addInventoryItem(equipement.Name, 1)
	for i, item := range p.inventaire {
		if item.Elem.Name == equipement.Name {
			if item.Elem.Level != equipement.Level {
				p.inventaire[i].Elem.Level = equipement.Level
				p.inventaire[i].Elem.BonusLife += 10
			}
		}
	}
	fmt.Println("You have unequipped ", equipement.Name, ", -", equipement.BonusLife, "PVMax")
	p.waitForEnter()
}

// Fonction pour vérifier si le personnage a une armure équipée
func (p *Personnage) accessEquipment() {
	if p.equipements.head.Name != "" {
		fmt.Println("1. Head:", p.equipements.head.Name, "+", p.equipements.head.BonusLife, "PVMax")
	}
	if p.equipements.body.Name != "" {
		fmt.Println("2. Body:", p.equipements.body.Name, "+", p.equipements.body.BonusLife, "PVMax")
	}
	if p.equipements.boots.Name != "" {
		fmt.Println("3. Boots:", p.equipements.boots.Name, "+", p.equipements.boots.BonusLife, "PVMax")
	}
	if p.equipements.head.Name == "" && p.equipements.body.Name == "" && p.equipements.boots.Name == "" {
		fmt.Println("You do not have any equipped items.")
	}
	fmt.Print("\nEnter the object number you want to unequip or press Enter to go back : ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		return
	}
	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		clearConsole()
		return
	}
	choix, err := strconv.Atoi(trimmedInput)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		return
	}
	switch {
	case choix == 1:
		if p.equipements.head.Name != "" {
			p.unequipArmor(p.equipements.head)
		} else {
			fmt.Println("You do not have any head equiped.")
			p.waitForEnter()
		}
	case choix == 2:
		if p.equipements.body.Name != "" {
			p.unequipArmor(p.equipements.body)
		} else {
			fmt.Println("You do not have any body equiped.")
			p.waitForEnter()
		}
	case choix == 3:
		if p.equipements.boots.Name != "" {
			p.unequipArmor(p.equipements.boots)
		} else {
			fmt.Println("You do not have any boots equiped.")
			p.waitForEnter()
		}
	default:
		fmt.Println("Input error.")
	}
}

// Fonction pour utiliser une potion de heal
func (p *Personnage) takePot(item Item) {
	if p.checkPV() < p.pvMax {
		fmt.Println("You drank a healing potion.")
		if len(item.Effect) > 0 {
			healAmount := item.Effect[0].Heal
			p.removeInventoryItem(item.Name, 1)
			p.addPV(healAmount)
			if p.checkPV() > p.pvMax {
				p.pv = p.pvMax
			}
			fmt.Println("Life points : ", p.checkPV(), "/", p.pvMax)
		} else {
			fmt.Println("The item has no healing effect.")
		}
	}
}

// Fonction pour utiliser une potion de poison
func (p *Personnage) poisonPot(item Item, m *monster) {
	fmt.Println("You throw a potion of poison.")
	p.removeInventoryItem(item.Name, 1)

	if len(item.Effect) > 0 {
		duration := item.Effect[0].Duration
		damages := item.Effect[0].Damages

		for i := 0; i < duration; i++ {
			time.Sleep(1 * time.Second)
			m.PV -= damages
			fmt.Println("Ennemy's Life points : ", m.PV, "/", m.PVMax)
		}
	} else {
		fmt.Println("The item has no poison effect.")
	}
}

// Fonction pour regarder l'inventaire
func (p *Personnage) accessInventory() {
	fmt.Println("Inventory: ")
	var count int = 1
	for _, item := range p.inventaire {
		fmt.Printf(" %v. %s (x%d)\n", count, item.Elem.Name, item.Quantity)
		count++
	}
	fmt.Print("Enter the object number to use or press Enter to go back : ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		return
	}

	trimmedInput := strings.TrimSpace(input)

	if trimmedInput == "" {
		clearConsole()
		return
	}

	choix, err := strconv.Atoi(trimmedInput)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		return
	}

	switch {
	case choix >= 1 && choix <= len(p.inventaire):
		if p.inventaire[choix-1].Elem.Type != "items" || p.inventaire[choix-1].Elem.Name == "Sac à dos" {
			//fmt.Println(p.inventaire[choix-1].Elem) <- test: affiche l'item sélectionné
			if p.inventaire[choix-1].Elem.Label == "potion_poison" {
				fmt.Println("You cannot use this item out of fight.")
				fmt.Println("Do you want to drank it and poisonned you ? Are you a masochist ?")
				p.waitForEnter()
			} else {
				p.useItem(p.inventaire[choix-1].Elem)
				p.waitForEnter()
			}
		} else if p.inventaire[choix-1].Elem.Type == "items" {
			fmt.Println("You cannot use this item.")
			p.waitForEnter()
		}
	default:
		fmt.Println("Input error.")
	}
}

// Fonction pour utiliser un item
func (p *Personnage) accessFightInventory(m *monster) bool {
	fmt.Println("Inventory : ")
	var count int = 1
	for _, item := range p.inventaire {
		if item.Elem.Type != "items" {
			fmt.Printf(" %v. %s (x%d)\n", count, item.Elem.Name, item.Quantity)
		}
		count++
	}

	fmt.Print("Enter the object number to use or press Enter to go back : ")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		return false
	}

	trimmedInput := strings.TrimSpace(input)

	if trimmedInput == "" {
		clearConsole()
		return false
	}

	choix, err := strconv.Atoi(trimmedInput)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		return false
	}

	switch {
	case choix >= 1 && choix <= len(p.inventaire):
		if p.inventaire[choix-1].Elem.Label == "potion_poison" {
			p.poisonPot(p.inventaire[choix-1].Elem, m)
		} else {
			p.useItem(p.inventaire[choix-1].Elem)
		}
	default:
		fmt.Println("Input error.")
	}
	return true
}

// Fonction pour augmenter la capacité de l'inventaire
func (p *Personnage) upgradeInventorySlot(itemName string) {
	if p.inventoryCapacity < 40 {
		if p.haveItem(itemName) {
			p.removeInventoryItem(itemName, 1)
			p.inventoryCapacity += 10
			fmt.Println("Your inventory capacity has been increased by 10.")
		} else {
			fmt.Println("You do not have this item in your inventory.")
			return
		}
	} else {
		fmt.Println("Your inventory capacity is already at maximum.")
		return
	}
}
