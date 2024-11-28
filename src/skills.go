package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

// Structure des skills
type Skills struct {
	Name     string
	Level    int
	ManaCost int `json:"ManaCost"`
	Attack   []struct {
		Name         string `json:"Name"`
		Damage       int    `json:"Damage"`
		Effect1      string `json:"Effect1"`
		EffectValue1 int    `json:"EffectValue1"`
		Effect2      string `json:"Effect2"`
		EffectValue2 int    `json:"EffectValue2"`
		Duration     int    `json:"Duration"`
	} `json:"Attack"`
}

// Structure du spellbook
type SpellBook struct {
	Name   string
	Skills []Skills
}

type SkillDatabase map[string][]Skills

// Fonction pour charger les skills
func loadSkillsFromFile() (SkillDatabase, error) {
	skillsData, err := ioutil.ReadFile("data/skills.json")
	if err != nil {
		return nil, err
	}

	var skillsMap SkillDatabase
	if err := json.Unmarshal(skillsData, &skillsMap); err != nil {
		return nil, err
	}

	return skillsMap, nil
}

// Fonction pour ajouter un skill à l'inventaire
func (p *Personnage) spellBook(spellbook SpellBook) {
	for _, spell := range spellbook.Skills {
		for _, skill := range p.skills {
			if skill.Name == spell.Name {
				return
			}
		}
		p.skills = append(p.skills, spell)
		fmt.Println("You have learned the skill :", spell.Name)
	}
	p.removeInventoryItem(spellbook.Name, 1)
	return
}

/*
func (p *Personnage) fuseSkill(skill1 SpellBook, skill2 SpellBook) {
	newSkill := Skills{
		Name:  skill1.Name + " " + skill2.Name,
		Level: 1,
	}

	for _, skill := range p.skills {
		if skill.Name == newSkill.Name {
			return
		}
	}

	p.skills = append(p.skills, newSkill)
	fmt.Println("You have learned the skill:", newSkill.Name)
	p.removeInventoryItem(skill1.Name, 1)
	p.removeInventoryItem(skill2.Name, 1)
	return
}*/

// Fonction pour augmenté le niveau des skills
func (p *Personnage) improveSkill(skillName string) {
	for i, skill := range p.skills {
		if p.skills[i].Level < 6 {
			if skill.Name == skillName {
				// calcule le prix en fonction du level
				var Level_Price float64
				if p.skills[i].Level == 1 {
					Level_Price = 1
					fmt.Println(Level_Price)
				} else if p.skills[i].Level == 5 {
					Level_Price = 33333
				} else {
					Level_Price = math.Pow(2, float64(p.skills[i].Level-1))
					fmt.Println(Level_Price)
				}
				// ajouter une condition qui vérifie si le joueur a assez d'argent
				if p.gold < 30*int(Level_Price) {
					fmt.Println("You do not have enough gold to improve this skill.")
					return
				} else {
					if p.skills[i].Level == 5 {
						fmt.Println("You TryHarder ! Don't you want to play another game or go take a bath ?")
						fmt.Println("Anyway, you can't improve this skill anymore... But thanks for your money !")
						p.gold -= 30*int(Level_Price) + 9
						return
					}
					p.gold -= 30 * int(Level_Price)
					p.skills[i].Level++
					fmt.Printf("You have improved the skill %s à Level %d for %v gold.\n", skillName, p.skills[i].Level, 30*int(Level_Price))
					return
				}
			}
		}
	}
	fmt.Printf("Skill %s was not found in your skills.\n", skillName)
}

// Fonction pour utiliser un skill
func (p *Personnage) UseSkill(selectedSkill Skills, m *monster) {
	baseDamage := 0
	ManaCost := 0

	skillsMap, err := loadSkillsFromFile()
	if err != nil {
		fmt.Println("Error loading skills:", err)
		return
	}

	skills, exists := skillsMap[p.classe]
	if !exists || len(skills) == 0 {
		fmt.Println("Skill not found")
		return
	}

	for _, s := range skills {
		if s.Name == selectedSkill.Name {
			baseDamage = s.Attack[0].Damage
			ManaCost = s.ManaCost
			break
		}
	}

	if baseDamage == 0 {
		fmt.Println("Skill level not found")
		return
	}

	damage := baseDamage
	for i := 1; i < selectedSkill.Level; i++ {
		damage = damage * 12 / 10
		ManaCost = ManaCost + 6
	}

	fmt.Printf("You used the skill : %s %v\n", selectedSkill.Name, selectedSkill.Level)
	fmt.Printf("You have inflicted %v damage on %s\n", damage, m.Name)
	p.SkillEffect(selectedSkill, m)
	p.removeMana(ManaCost)
	m.PV -= damage
}

// Fonction pour les effets des skills
func (p *Personnage) SkillEffect(selectedSkill Skills, m *monster) {
	if selectedSkill.Attack[0].Effect1 == "Poison" {
		damages := selectedSkill.Attack[0].EffectValue1
		for i := 0; i < selectedSkill.Level; i++ {
			damages = damages * 12 / 10
		}
		fmt.Printf("You have poisoned %s for %v seconds\n", m.Name, selectedSkill.Attack[0].Duration)
		duration := selectedSkill.Attack[0].Duration
		for i := 0; i < duration; i++ {
			m.PV -= damages
			fmt.Printf("%s takes %v damages from poison\n", m.Name, damages)
			time.Sleep(1 * time.Second)
		}
	} else if selectedSkill.Attack[0].Effect2 == "Poison" {
		damages := selectedSkill.Attack[0].EffectValue2
		for i := 0; i < selectedSkill.Level; i++ {
			damages = damages * 12 / 10
		}
		fmt.Printf("You have poisoned %s for %v seconds\n", m.Name, selectedSkill.Attack[0].Duration)
		duration := selectedSkill.Attack[0].Duration
		for i := 0; i < duration; i++ {
			m.PV -= damages
			fmt.Printf("%s takes %v damages from poison\n", m.Name, damages)
			time.Sleep(1 * time.Second)
		}
	}
	if selectedSkill.Attack[0].Effect1 == "Fire" {
		damages := selectedSkill.Attack[0].EffectValue1
		for i := 0; i < selectedSkill.Level; i++ {
			damages = damages * 12 / 10
		}
		fmt.Println("You have burned", m.Name, "he will take damages at the begining of his turn until the end of the fight")
		m.Status = "Burned"
		m.StatusEffect = damages
	} else if selectedSkill.Attack[0].Effect2 == "Fire" {
		damages := selectedSkill.Attack[0].EffectValue2
		for i := 0; i < selectedSkill.Level; i++ {
			damages = damages * 12 / 10
		}
		fmt.Println("You have burned", m.Name, "he will take damages at the begining of his turn until the end of the fight")
		m.Status = "Burned"
		m.StatusEffect = damages
	}
	if selectedSkill.Attack[0].Effect1 == "Shield" {
		Value := selectedSkill.Attack[0].EffectValue1
		for i := 0; i < selectedSkill.Level; i++ {
			Value = Value * 12 / 10
		}
		fmt.Println("You have shielded yourself")
		p.shield += Value
	} else if selectedSkill.Attack[0].Effect2 == "Shield" {
		Value := selectedSkill.Attack[0].EffectValue2
		for i := 0; i < selectedSkill.Level; i++ {
			Value = Value * 12 / 10
		}
		fmt.Println("You have shielded yourself")
		p.shield += Value
	}
}
