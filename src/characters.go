package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// Structure du personnage
type Personnage struct {
	nom               string
	classe            string
	niveau            int
	xp                int
	neededXp          int
	gold              int
	pv                int
	pvMax             int
	Attack            int
	shield            int
	Mana              int
	inventoryCapacity int
	inventaire        []Inventory
	skills            []Skills
	equipements       armor
	Death             int
}

// Structure des équipements
type armor struct {
	head  Item
	body  Item
	boots Item
}

// Structure sauvegarde du personnage
type SavedCharacter struct {
	Nom               string
	Classe            string
	Niveau            int
	XP                int
	NeededXP          int
	Gold              int
	PV                int
	PVMax             int
	Attack            int
	Shield            int
	Mana              int
	InventoryCapacity int
	Inventaire        []Inventory
	Skills            []Skills
	Equipements       armor
	Death             int
	EquippedHead  Item `json:"equipped_head"`
    EquippedBody  Item `json:"equipped_body"`
    EquippedBoots Item `json:"equipped_boots"`
}

// Initialisation du personnage
func Init(nom string, classe string, niveau int, gold int, pv int, pvMax int, attack int, shield int, skills []Skills, inventoryCapacity int, inventaireInitial []Inventory) Personnage {
	return Personnage{
		nom:               nom,
		classe:            classe,
		niveau:            niveau,
		xp:                0,
		neededXp:          100,
		gold:              gold,
		pv:                pv,
		pvMax:             pvMax,
		Attack:            attack,
		shield:            shield,
		Mana:              100,
		inventoryCapacity: inventoryCapacity,
		inventaire:        inventaireInitial,
		skills:            skills,
		Death:             0,
	}
}

// Fonction de sauvegarde des données du personnage
func (p *Personnage) saveCharacterToFile() error {
	savedChar := SavedCharacter{
		Nom:               p.nom,
		Classe:            p.classe,
		Niveau:            p.niveau,
		XP:                p.xp,
		NeededXP:          p.neededXp,
		Gold:              p.gold,
		PV:                p.pv,
		PVMax:             p.pvMax,
		Attack:            p.Attack,
		Shield:            p.shield,
		InventoryCapacity: p.inventoryCapacity,
		Inventaire:        p.inventaire,
		Skills:            p.skills,
		Equipements:       p.equipements,
		Death:             p.Death,
		EquippedHead:  p.equipements.head,
        EquippedBody:  p.equipements.body,
        EquippedBoots: p.equipements.boots,
	}

	data, err := json.Marshal(savedChar)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/savePersonnage.json", data, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Character data saved successfully.")
	return nil
}

// Fonction de chargement des données du personnage
func loadCharacterFromFile(filename string) (*Personnage, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var savedChar SavedCharacter
	if err := json.Unmarshal(data, &savedChar); err != nil {
		return nil, err
	}

	perso := &Personnage{
		nom:               savedChar.Nom,
		classe:            savedChar.Classe,
		niveau:            savedChar.Niveau,
		xp:                savedChar.XP,
		neededXp:          savedChar.NeededXP,
		gold:              savedChar.Gold,
		pv:                savedChar.PV,
		pvMax:             savedChar.PVMax,
		Attack:            savedChar.Attack,
		shield:            savedChar.Shield,
		inventoryCapacity: savedChar.InventoryCapacity,
		inventaire:        savedChar.Inventaire,
		skills:            savedChar.Skills,
		equipements: armor{
            head:  savedChar.EquippedHead,
            body:  savedChar.EquippedBody,
            boots: savedChar.EquippedBoots,
        },
		Death:             savedChar.Death,
	}

	return perso, nil
}

// Fonction d'affichage des informations du personnage
func (p *Personnage) displayInfo() {
	fmt.Println("Name: " + p.nom)
	fmt.Println("Class: " + p.classe)
	fmt.Println("Level: ", p.checkLevel())
	fmt.Println("XP: ", p.xp, "/", p.neededXp)
	fmt.Println("Gold: ", p.checkGold())
	fmt.Println("Inventory Capacity: ", p.inventoryCapacity, "/ 40")
	fmt.Println("Life point: ", p.checkPV(), "/", p.pvMax)
	fmt.Println("Attack: ", p.Attack)
	fmt.Println("Shield: ", p.checkShield())
	fmt.Println("Mana: ", p.Mana)
	fmt.Println("Number of death: ", p.Death)
}

// Fonction de vérification du nom du personnage
func isValidName(name string) bool {
	if len(name) == 0 {
		return false
	}

	if !unicode.IsUpper([]rune(name)[0]) {
		return false
	}

	for _, char := range name[1:] {
		if !unicode.IsLetter(char) || !unicode.IsLower(char) {
			return false
		}
	}

	return true
}

// Fonction de vérification du nom du personnage
func getValidName() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please choose your character's name.")
		fmt.Println("")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		name = strings.Title(strings.ToLower(name))
		if isValidName(name) {
			return name
		} else {
			fmt.Println("The name must begin with an uppercase letter followed by lowercase letters, without numbers or spaces.")
		}
	}
}

// Fonction de vérification de la classe du personnage
func getValidClass() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Please write your character's class.")
		fmt.Println("Classes available : Guerrier, Archer, Sorcier")
		fmt.Println("")
		class, _ := reader.ReadString('\n')
		class = strings.TrimSpace(class)
		class = strings.Title(strings.ToLower(class))
		if isValidName(class) {
			return class
		} else {
			fmt.Println("The class must begin with an uppercase letter followed by lowercase letters, without numbers or spaces.")
		}
	}
}

// Fonction pour ajouter de l'xp au personnage
func (p *Personnage) addXP(xp int) {
	p.xp += xp
}

// Fonction pour ajouter du mana au personnage
func (p *Personnage) addMana(mana int) {
	p.Mana += mana
}

// Fonction pour enlever de l'xp au personnage
func (p *Personnage) removeXP(xp int) {
	p.xp -= xp
}

// Fonction pour enlever du mana au personnage
func (p *Personnage) removeMana(mana int) {
	p.Mana -= mana
}

// Fonction d'affichage des skills du personnage
func (p *Personnage) checkSkills() {
	fmt.Println("Skills: ")
	for _, skill := range p.skills {
		fmt.Printf(" - %s (Level %d)\n", skill.Name, skill.Level)
	}
}

// Fonction pour vérifier si le personnage est mort
func (p *Personnage) dead() bool {
	if p.checkPV() <= 0 {
		fmt.Println("You died !")
		fmt.Println("You will be resurrected with half your life and 50 mana point in addition.")
		if p.Death == 0 {
			fmt.Println("Be careful, next times you'll die, you'll lose 10%", "of your gold.")
		} else if p.Death > 0 {
			fmt.Println("You lost 10%", "of your gold.")
			p.gold = p.gold - (p.gold / 10)
		}
		p.Death++
		p.pv = p.pvMax / 2
		p.shield = 0
		p.Mana += 50
		return true
	}
	return false
}

// Fonction pour check du gold
func (p *Personnage) checkGold() int {
	return p.gold
}

// Fonction pour check des pv
func (p *Personnage) checkPV() int {
	return p.pv
}

// Fonction pour check le shield
func (p *Personnage) checkShield() int {
	return p.shield
}

// Fonction pour check le level
func (p *Personnage) checkLevel() int {
	return p.niveau
}

// Fonction pour ajouter des pv
func (p *Personnage) addPV(pv int) {
	p.pv += pv
}

// Fonction pour ajouter du shield
func (p *Personnage) addShield(shield int) {
	p.shield += shield
}

// Fonction pour ajouter du gold
func (p *Personnage) addGold(gold int) {
	p.gold += gold
}

// Fonction pour enlever des pv
func (p *Personnage) removePV(pv int) {
	p.pv -= pv
}

// Fonction pour enlever du shield
func (p *Personnage) removeShield(shield int) {
	p.shield -= shield
}

// Fonction pour enlever du gold
func (p *Personnage) removeGold(gold int) {
	p.gold -= gold
}

// Fonction pour créer le personnage
func CreateChar() (string, string, int) {
	var name string
	var classe string
	var pvMaxPerso int

	name = getValidName()
	classe = getValidClass()

	switch classe {
	case "Guerrier":
		pvMaxPerso = 120
	case "Archer":
		pvMaxPerso = 100
	case "Sorcier":
		pvMaxPerso = 80
	}

	fmt.Println("You have chosen the class:", classe)
	fmt.Println("")
	return name, classe, pvMaxPerso
}

// Fonction pour level up le personnage
func (p *Personnage) LevelUp() {
	if p.xp >= p.neededXp {
		p.niveau++
		p.xp -= p.neededXp
		p.neededXp *= (15 / 10)
		p.pvMax += 10
		p.pv = p.pvMax
		p.Attack += 5
		p.shield += 5
		fmt.Println("You have gained a level!")
		fmt.Println("Your health points have increased by 10 points and have been restored.")
		fmt.Println("Your attack has increased by 5 points.")
		fmt.Println("Your defense has increased by 5 points.")
		fmt.Println("")
		if p.xp >= p.neededXp {
			p.LevelUp()
		}
	}
}
