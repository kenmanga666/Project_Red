package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Structure du forgerondata
type ForgeronData struct {
	Data []Item `json:"Forgeron"`
}

// Fonction back forgeron
func (p *Personnage) waitForEnterForgeron() {
	fmt.Println("Press Enter to go back...")
	fmt.Scanln()
	clearConsole()
	p.forgeron()
}

// Fonction menu forgeron
func (p *Personnage) forgeron() {
	for {
		fmt.Println("What do you want to do ?")

		fmt.Println("1. Craft items")
		fmt.Println("2. Improve stuff")
		fmt.Println("3. Dismantle items")
		fmt.Println("0. Back")
		fmt.Println("")

		var choice int
		fmt.Print("Enter your choice : ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Input error. Please enter a valid choice.")
			continue
		}

		switch choice {
		case 1:
			clearConsole()
			p.craftMenu()
		case 2:
			clearConsole()
			p.upgradeMenu()
		case 3:
			clearConsole()
			p.dismantleMenu()
		case 0:
			clearConsole()
			p.menu()
		default:
			fmt.Println("Input error. Please enter a valid choice.")
			continue
		}
	}
}

// Fonction menu craft
func (p *Personnage) craftMenu() {
	for {
		data, err := ioutil.ReadFile("data/forgeron.json")
		if err != nil {
			fmt.Println(err)
		}

		var items ForgeronData
		if err := json.Unmarshal(data, &items); err != nil {
			fmt.Println(err)
		}

		for _, item := range items.Data {
			fmt.Printf("%s\n", item.Label)
		}
		fmt.Println("0. Back")
		fmt.Println("")

		var choixForgeron int
		fmt.Print("Enter your choice : ")
		_, err = fmt.Scanln(&choixForgeron)
		if err != nil {
			fmt.Println("Input error. Please enter a valid choice.")
			continue
		}

		switch choixForgeron {
		case 1, 2, 3:
			clearConsole()
			p.itemInfo(items.Data[choixForgeron-1])
		case 0:
			clearConsole()
			p.forgeron()
		default:
			fmt.Println("Input error. Please enter a valid choice.")
		}
	}
}

// Fonction info item
func (p *Personnage) itemInfo(item Item) {
	fmt.Println("Materials needed :")
	for _, craft := range item.Craft {
		fmt.Printf("- %s: %d\n", craft.Name, craft.Quantity)
	}
	fmt.Println("")
	fmt.Println("1. Make")
	fmt.Println("0. Back")

	var choix int
	fmt.Print("Enter your choice : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.itemInfo(item)
	}

	switch choix {
	case 1:
		clearConsole()
		p.craftItem(item)
	case 0:
		clearConsole()
		p.forgeron()
	default:
		fmt.Println("Invalid choice. Please choose a valid option.")
		p.itemInfo(item)
	}
}

// Fonction craft item
func (p *Personnage) craftItem(item Item) {
	hasMaterials := false
	for _, craft := range item.Craft {
		if craft.Name == "Pièces d'or" {
			continue
		}

		if p.haveItem(craft.Name) && p.getInventoryItemQuantity(craft.Name) >= craft.Quantity {
			hasMaterials = true
		} else {
			hasMaterials = false
			break
		}
	}

	if hasMaterials && p.checkGold() >= 5 {
		for _, craft := range item.Craft {
			if craft.Name == "Pièces d'or" {
				continue
			}
			p.removeInventoryItem(craft.Name, craft.Quantity)
		}
		p.removeGold(5)

		p.addInventoryItem(item.Name, 1)

		fmt.Printf("You made a %s.\n", item.Name)
		p.waitForEnterForgeron()
	} else {
		clearConsole()
		fmt.Println("You do not have the required materials or enough gold.")
		p.waitForEnterForgeron()
	}
}

// Fonction upgrade menu
func (p *Personnage) upgradeMenu() {
    for {
		clearConsole()
		fmt.Println("Equipment that can be improved:")

        var upgradableItems []Item
		for _, inventoryItem := range p.inventaire {
			if inventoryItem.Elem.Type == "equipement" {
				upgradableItems = append(upgradableItems, inventoryItem.Elem)
			}
		}

        if len(upgradableItems) == 0 {
			fmt.Println("You do not have any upgradable equipment.")
			fmt.Println("0. Back")
			fmt.Print("\nEnter your choice : ")
		} else {
			for i, item := range upgradableItems {
				fmt.Printf("%d. %s (Level %d)\n", i+1, item.Name, item.Level)
			}
			fmt.Println("0. Back")
			fmt.Print("\nEnter your choice : ")
		}

        var choixUpgrade int
		_, err := fmt.Scanln(&choixUpgrade)
		if err != nil {
			fmt.Println("Input error. Please enter a valid choice.")
			continue
		}

        switch choixUpgrade {
		case 0:
			clearConsole()
			p.forgeron()
		default:
			if choixUpgrade >= 1 && choixUpgrade <= len(upgradableItems) {
				selectedItem := upgradableItems[choixUpgrade-1]

				cost := 0
				if selectedItem.Level == 1 {
					cost = 15
				} else if selectedItem.Level == 2 {
					cost = 25
				} else if selectedItem.Level == 3 {
					cost = 1001
				}

				if p.checkGold() >= cost {
					clearConsole()
					p.upgradeItem(selectedItem, cost)
				} else {
					clearConsole()
					fmt.Println("You do not have enough gold to upgrade this equipment.")
					p.waitForEnterForgeron()
				}
			} else {
				fmt.Println("Invalid choice. Please choose a valid option.")
			}
		}
    }
}

// Fonction upgrade item
func (p *Personnage) upgradeItem(item Item, cost int) {
    fmt.Printf("Equipment Improvement : %s (Level %d -> Level %d)\n", item.Name, item.Level, item.Level+1)
	fmt.Printf("Cost : %d gold\n", cost)
	fmt.Println("1. Confirm improvement")
	fmt.Println("2. Cancel")

	var choix int
	fmt.Print("Enter your choice : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.upgradeMenu()
		return
	}

	switch choix {
	case 1:
		clearConsole()
		if item.Level < 3 {
			p.removeGold(cost)
		
			for i, invItem := range p.inventaire {
				if invItem.Elem.Name == item.Name {
					p.inventaire[i].Elem.Level++
                    p.inventaire[i].Elem.BonusLife += 10
					break
				}
			}
			fmt.Printf("%s has been improved to the level %d.\n", item.Name, item.Level+1)
			p.waitForEnterForgeron()
		} else if item.Level == 3 {
            fmt.Println("You Sucker ! Why do you want to Upgrade your equipment more ?")
			fmt.Println("This game is as hard as that ??")
			fmt.Println("Anyway, heres you're equipment but don't try to look at it... Also thanks for your money !")
			p.removeGold(cost)
			p.removeInventoryItem(item.Name, 1)
            if item.Name == "Chapeau de l'aventurier" {
				p.addInventoryItem("Sucker's Hat", 1)
			} else if item.Name == "Tunique de l'aventurier" {
				p.addInventoryItem("Sucker's Tunic", 1)
			} else if item.Name == "Bottes de l'aventurier" {
				p.addInventoryItem("Sucker's Boots", 1)
			}
			p.waitForEnterForgeron()
        } else {
			fmt.Println("This equipment is already at maximum level.")
            fmt.Println("Don't you have something else to do ?")
			return
		}
    case 2:
		clearConsole()
		p.upgradeMenu()
	default:
		fmt.Println("Invalid choice. Please choose a valid option.")
		p.upgradeItem(item, cost)
	}
}

// Fonction dismantle menu
func (p *Personnage) dismantleMenu() {
    for {
        clearConsole()
		fmt.Println("Inventory:")

		var itemsToDismantle []Item
		for _, item := range p.inventaire {
			if item.Elem.Type != "items" {
				itemsToDismantle = append(itemsToDismantle, item.Elem)
			}
		}

        if len(itemsToDismantle) == 0 {
			fmt.Println("You do not have any items available for dismantling.")
			fmt.Println("0. Back")
			fmt.Print("\nEnter your choice : ")
		} else {
			for i, item := range itemsToDismantle {
				fmt.Printf("%d. %s (x%d)\n", i+1, item.Name, p.getInventoryItemQuantity(item.Name))
			}
			fmt.Println("0. Back")
			fmt.Print("\nEnter your choice : ")
		}

        var choixDismantle int
		_, err := fmt.Scanln(&choixDismantle)
		if err != nil {
			fmt.Println("Input error. Please enter a valid choice.")
			continue
		}

        switch choixDismantle {
		case 0:
			clearConsole()
			p.forgeron()
		default:
			if choixDismantle >= 1 && choixDismantle <= len(itemsToDismantle) {
				selectedItem := itemsToDismantle[choixDismantle-1]

				forgeronData, err := ioutil.ReadFile("data/forgeron.json")
				if err != nil {
					fmt.Println(err)
					continue
				}

				var forgeronItems ForgeronData
				if err := json.Unmarshal(forgeronData, &forgeronItems); err != nil {
					fmt.Println(err)
					continue
				}

				for _, forgeronItem := range forgeronItems.Data {
					if forgeronItem.Name == selectedItem.Name {
						clearConsole()
						p.dismantleItem(forgeronItem)
						break
					}
				}
			} else {
				fmt.Println("Input error. Please enter a valid choice.")
			}
		}
    }
}

// Fonction dismantle item
func (p *Personnage) dismantleItem(item Item) {
    clearConsole()
	fmt.Printf("Dismantling the element : %s\n", item.Name)

	fmt.Println("Materials obtained :")
	for _, craft := range item.Craft {
		fmt.Printf("- %s: %d\n", craft.Name, craft.Quantity)
	}

    fmt.Println("")
	fmt.Println("1. Confirm dismantling")
	fmt.Println("2. Cancel")

	var choix int
	fmt.Print("Enter your choice : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.dismantleMenu()
		return
	}

    switch choix {
	case 1:
		clearConsole()
		p.removeInventoryItem(item.Name, 1)
	
		for _, craft := range item.Craft {
			if craft.Name != "Pièces d'or" {
				p.addInventoryItem(craft.Name, craft.Quantity)
				p.addGold(5)
			}
		}
        fmt.Printf("You have dismantled %s and obtained the corresponding materials.\n", item.Name)
		p.waitForEnterForgeron()
	case 2:
		clearConsole()
		p.dismantleMenu()
	default:
		fmt.Println("Invalid choice. Please choose a valid option.")
		p.dismantleItem(item)
	}
}
