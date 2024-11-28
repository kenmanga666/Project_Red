package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Structure des items
type Item struct {
	Name      string `json:"Name"`
	Label     string `json:"Label"`
	Type      string `json:"Type"`
	Price     int    `json:"Price"`
	SubType   string `json:"SubType"`
	BonusLife int    `json:"BonusLife"`
	Level     int      `json:"Level"`
	Effect    []struct {
		Name     string `json:"Name"`
		Type     string `json:"Type"`
		Heal     int    `json:"Heal"`
		Damages  int    `json:"Damages"`
		Duration int    `json:"Duration"`
	} `json:"Effect"`
	Craft []struct {
		Name     string `json:"Name"`
		Quantity int    `json:"Quantity"`
	} `json:"Craft"`
}

// Fonction pour ajouter un item à l'inventaire
func (p *Personnage) addInventoryItem(itemName string, quantity int) {
	jsonData, err := ioutil.ReadFile("data/items.json")
	if err != nil {
		fmt.Println("Error reading JSON file :", err)
		return
	}
	var items []Item
	if err := json.Unmarshal(jsonData, &items); err != nil {
		fmt.Println("Error parsing JSON file :", err)
		return
	}
	var itemExists bool
	var itemToAdd Item
	for _, item := range items {
		if item.Name == itemName {
			itemExists = true
			itemToAdd = item
			break
		}
	}
	if !itemExists {
		fmt.Printf("%s does not exist in items.json file.\n", itemName)
		return
	}
	if len(p.inventaire) >= p.inventoryCapacity {
		fmt.Println("The inventory is full. You cannot add more objects.")
		return
	}
	for i, item := range p.inventaire {
		if item.Elem.Name == itemName {
			if item.Quantity >= 10 {
				fmt.Printf("You have reached the limit of 10 for %s in your inventory.\n", itemName)
				return
			}
			p.inventaire[i].Quantity++
			fmt.Printf("You have added %s to your inventory. (x%d)\n", itemName, p.inventaire[i].Quantity)
			return
		}
	}
	newItem := Inventory{Elem: itemToAdd, Quantity: 1}
	p.inventaire = append(p.inventaire, newItem)
	fmt.Printf("You have added %s to your inventory. (x1)\n", itemName)
}

// Fonction pour enlever un item de l'inventaire
func (p *Personnage) removeInventoryItem(itemName string, quantity int) {
	for i, item := range p.inventaire {
		if item.Elem.Name == itemName {
			p.inventaire[i].Quantity -= quantity
			if p.inventaire[i].Quantity <= 0 {
				p.inventaire = append(p.inventaire[:i], p.inventaire[i+1:]...)
			}
			return
		}
	}
}

// Fonction pour afficher la quantité des items
func (p *Personnage) getInventoryItemQuantity(itemName string) int {
	for _, inventoryItem := range p.inventaire {
		if inventoryItem.Elem.Name == itemName {
			return inventoryItem.Quantity
		}
	}
	return 0
}

// Fonction pour vérifier si le player a l'item
func (p *Personnage) haveItem(itemName string) bool {
	jsonData, err := ioutil.ReadFile("data/items.json")
	if err != nil {
		fmt.Println("Error reading JSON file :", err)
		return false
	}
	var items []Item
	if err := json.Unmarshal(jsonData, &items); err != nil {
		fmt.Println("Error parsing JSON file :", err)
		return false
	}
	for _, item := range items {
		if item.Name == itemName {
			for _, invItem := range p.inventaire {
				if invItem.Elem.Name == itemName {
					return true
				}
			}
			return false
		}
	}
	return false
}

// Fonction pour utiliser les items
func (p *Personnage) useItem(item Item) {
	if item.Type == "potion" {
		if item.Label == "potion_soin" {
			p.takePot(item)
		}
	} else if item.Type == "equipement" {
		p.equipArmor(item)
	} else if item.Name == "Sac à dos" {
		p.upgradeInventorySlot(item.Name)
	}
}
