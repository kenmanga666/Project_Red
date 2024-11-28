package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

// Structure du magasin
type StoreData struct {
	Data []Item `json:""`
}

// Fonction back magasin
func (p *Personnage) waitForEnterStore() {
	fmt.Println("Press Enter to go back...")
	fmt.Scanln()
	clearConsole()
	p.enterShop()
}

// Fonction menu magasin
func (p *Personnage) enterShop() {
	fmt.Println("Welcome to the shop !")
	fmt.Println("What do you want to do ?")
	fmt.Println("1. Buy")
	fmt.Println("2. Sell")
	fmt.Println("3. Improve a skill")
	fmt.Println("0. Back")
	fmt.Println("")

	var choix int
	fmt.Print("Enter your choice : ")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		clearConsole()
		p.enterShop()
	}

	switch choix {
	case 1:
		clearConsole()
		p.buyMenu()
		p.waitForEnter()
		p.menu()
	case 2:
		clearConsole()
		p.sellMenu()
		p.waitForEnter()
		p.menu()
	case 3:
		clearConsole()
		p.skillImproveMenu()
		p.waitForEnter()
		p.menu()
	case 0:
		clearConsole()
		p.menu()
	default:
		fmt.Println("Choix invalide. Veuillez choisir une option valide.")
		p.enterShop()
	}
}

// Fonction chargement des items du magasin
func (p *Personnage) sellMenu() {
	jsonData, err := ioutil.ReadFile("data/items.json")
	if err != nil {
		fmt.Println("Error reading JSON file :", err)
		return
	}

	var allItems []Item
	if err := json.Unmarshal(jsonData, &allItems); err != nil {
		fmt.Println("Error parsing JSON file :", err)
		return
	}

	fmt.Println("What item do you want to sell?")

	var itemsToSell []Item
	for _, item := range allItems {
		if item.Type == "items" || item.Type == "potion" {
			for _, inventoryItem := range p.inventaire {
				if item.Name == inventoryItem.Elem.Name {
					itemsToSell = append(itemsToSell, item)
				}
			}
		}
	}

	if len(itemsToSell) == 0 {
		fmt.Println("You do not have any items to sell.")
		fmt.Println("0. Back")
		fmt.Print("\nEnter your choice : ")
	} else {
		for j, item := range itemsToSell {
			fmt.Printf("%d. %s (x%d). Selling price: %d\n", j+1, item.Name, p.getInventoryItemQuantity(item.Name), item.Price/2)
		}
		fmt.Println("0. Back")
		fmt.Print("\nEnter your choice : ")
	}

	var choix int

	fmt.Print("Enter your choice :")
	_, err = fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.sellMenu()
	}

	if choix == 0 {
		clearConsole()
		p.enterShop()
	} else if choix >= 1 && choix <= len(p.inventaire) {
		itemIndex := choix - 1
		p.sellItem(p.inventaire[itemIndex].Elem.Name, allItems)
	} else {
		fmt.Println("Invalid choice. Please choose a valid option.")
		p.sellMenu()
	}
}

// Fonction pour vendre un item
func (p *Personnage) sellItem(itemName string, allItems []Item) {
	var itemToSell Item
	for _, item := range allItems {
		if item.Name == itemName {
			itemToSell = item
			break
		}
	}

	if itemToSell.Name == "" {
		fmt.Printf("Item %s does not exist in the JSON file.\n", itemName)
		return
	}

	if p.haveItem(itemName) {
		for i, item := range p.inventaire {
			if item.Elem.Name == itemName {
				sellPrice := itemToSell.Price / 2
				p.inventaire[i].Quantity--
				if p.inventaire[i].Quantity <= 0 {
					p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
				}
				p.addGold(sellPrice)
				clearConsole()
				fmt.Printf("You sold %s for %d gold.\n", itemName, sellPrice)
				p.waitForEnterStore()
			}
		}
	} else {
		fmt.Printf("You do not have %s in your inventory.\n", itemName)
	}
}

// Fonction menu acheter
func (p *Personnage) buyMenu() {
	fmt.Println("What item do you want to buy ?")

	jsonData, err := ioutil.ReadFile("data/items.json")
	if err != nil {
		fmt.Println("Error reading JSON file :", err)
		return
	}

	var allItems []Item
	if err := json.Unmarshal(jsonData, &allItems); err != nil {
		fmt.Println("Error parsing JSON file :", err)
		return
	}

	itemIndex := 0

	for _, item := range allItems {
		if item.Type == "items" || item.Type == "potion" {
			itemIndex++
			fmt.Printf("%d. %s. Price : %d\n", itemIndex, item.Name, item.Price)
		}
	}

	var selectedItemIndex int
	fmt.Println("0. Back")
	fmt.Println("")
	fmt.Print("Choose an item number to purchase :")
	_, err = fmt.Scan(&selectedItemIndex)
	if err != nil {
		fmt.Println("Error reading user input :", err)
		return
	}

	if selectedItemIndex == 0 {
		clearConsole()
		p.enterShop()
	} else if selectedItemIndex < 1 || selectedItemIndex > itemIndex {
		fmt.Println("Invalid object index.")
		p.waitForEnterStore()
	}

	selectedIdx := -1
	count := 0
	for idx, item := range allItems {
		if item.Type == "items" || item.Type == "potion" {
			count++
			if count == selectedItemIndex {
				selectedIdx = idx
				break
			}
		}
	}

	if selectedIdx != -1 {
		p.buyItem(allItems[selectedIdx])
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	} else {
		clearConsole()
		fmt.Println("The selected object index was not found.")
		p.waitForEnterStore()
	}
}

// Fonction pour acheter un item
func (p *Personnage) buyItem(item Item) {
	if len(p.inventaire) >= p.inventoryCapacity {
		fmt.Println("The inventory is full. You cannot add more objects.")
		return
	}
	if p.gold < item.Price {
		fmt.Println("You do not have enough money to buy this item.")
		return
	}
	fmt.Println("Item Price :", item.Price, "gold.")
	p.gold -= item.Price
	p.addInventoryItem(item.Name, 1)
	fmt.Printf("You bought %s for %d gold.\n", item.Name, item.Price)
	p.waitForEnterStore()
}

// Fonction menu améliorer skill
func (p *Personnage) skillImproveMenu() {
	fmt.Println("What skill do you want to improve?")
	for i, skill := range p.skills {
		var Level_Price float64
		var Troll_Bonus int
		if p.skills[i].Level == 1 {
			Level_Price = 1
		} else if p.skills[i].Level == 5 {
			Level_Price = 33333
			Troll_Bonus = 9
		} else {
			Level_Price = math.Pow(2, float64(p.skills[i].Level-1))
		}
		fmt.Printf("%d. %s (Level %d) Price: %v gold\n", i+1, skill.Name, skill.Level, 30*int(Level_Price)+Troll_Bonus)
	}
	fmt.Println("0. Back")
	fmt.Println("")

	var choix int
	fmt.Print("Enter your choice :")
	_, err := fmt.Scanln(&choix)
	if err != nil {
		fmt.Println("Input error.")
		p.skillImproveMenu()
	}

	if choix == 0 {
		clearConsole()
		p.enterShop()
	} else if choix >= 1 && choix <= len(p.skills) {
		skillIndex := choix - 1
		p.improveSkill(p.skills[skillIndex].Name)
		p.waitForEnterStore()
	} else {
		fmt.Println("Invalid choice. Please choose a valid option.")
		p.skillImproveMenu()
	}
}
