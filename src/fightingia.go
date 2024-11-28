package main

import (
	"fmt"
	"math/rand"
)

// Structure du monstre
type monster struct {
	Name         string
	PVMax        int
	PV           int
	Attack       int
	Status       string
	StatusEffect int
}

// Fonction initalisation du monstre
func (m *monster) InitGoblin() {
	m.Name = "Goblin"
	m.PVMax = 40
	m.PV = m.PVMax
	m.Attack = 5
}

// Fonction initalisation du monstre
func (m *monster) InitOrc() {
	m.Name = "Orc"
	m.PVMax = 60
	m.PV = m.PVMax
	m.Attack = 10
}

// Fonction initalisation du boss
func (m *monster) InitSecretBoss() {
	m.Name = "World Eater"
	m.PVMax = 500
	m.PV = m.PVMax
	m.Attack = 30
}

// Fonction pour l'initialisation du monstre
func (m *monster) InitMonster(name string) {
	fmt.Println(name)
	if name == "Goblin" {
		m.InitGoblin()
	} else if name == "Orc" {
		m.InitOrc()
	} else if name == "SecretBoss" {
		m.InitSecretBoss()
	}
}

// Fonction pour le pattern du monstre
func (m *monster) GoblinPattern(tour int, p *Personnage) {
	if tour%3 == 0 {
		fmt.Println("Critical Hit !", m.Name, "deals to", p.nom, m.Attack*2, "damages")
		if p.shield == 0 {
			p.removePV(m.Attack * 2)
		} else if p.shield > 0 && p.shield <= m.Attack*2 {
			dgttopv := (m.Attack * 2) - p.shield
			p.removePV(dgttopv)
			p.shield = 0
		} else {
			p.removeShield(m.Attack * 2)
		}
	} else {
		fmt.Println(m.Name, "deals to", p.nom, m.Attack, "damages")
		if p.shield == 0 {
			p.removePV(m.Attack)
		} else if p.shield > 0 && p.shield <= m.Attack {
			dgttopv := (m.Attack) - p.shield
			p.removePV(dgttopv)
			p.shield = 0
		} else {
			p.removeShield(m.Attack)
		}
	}
	fmt.Println("your remaining Shield", p.checkShield())
	fmt.Println("your remaining PVs", p.checkPV(), "/", p.pvMax)
}

// Fonction pour le pattern du monstre
func (m *monster) OrcPattern(tour int, p *Personnage) {
	if tour%3 == 0 {
		fmt.Println(m.Name, "is enraged !", m.Name, "Attack multiplied by 3 for this turn!")
		fmt.Println(m.Name, "deals to", p.nom, m.Attack*3, "damages")
		if p.shield == 0 {
			p.removePV(m.Attack * 3)
		} else if p.shield > 0 && p.shield <= m.Attack*3 {
			dgttopv := (m.Attack * 3) - p.shield
			p.removePV(dgttopv)
			p.shield = 0
		} else {
			p.removeShield(m.Attack * 3)
		}
	} else {
		fmt.Println(m.Name, "deals to", p.nom, m.Attack, "damages")
		if p.shield == 0 {
			p.removePV(m.Attack)
		} else if p.shield > 0 && p.shield <= m.Attack {
			dgttopv := (m.Attack) - p.shield
			p.removePV(dgttopv)
			p.shield = 0
		} else {
			p.removeShield(m.Attack)
		}
	}
	fmt.Println("your remaining Shield", p.checkShield())
	fmt.Println("your remaining PVs", p.checkPV(), "/", p.pvMax)
}

// Fonction pour le pattern du monstre
func (m *monster) SecretBossPattern(tour int, p *Personnage) {
	CriticalChance := rand.Intn(10)
	CriticalDamage := rand.Intn(10)
	fmt.Println("Boss's Critical Chance: ", CriticalChance)
	fmt.Println("Boss's Critical Damages: ", CriticalDamage)
	if CriticalChance == 0 {
		fmt.Println(m.Name, "deals to", p.nom, m.Attack, "damages")
		if p.shield == 0 {
			p.removePV(m.Attack)
			m.PV += m.Attack * (20 / 100)
		} else if p.shield > 0 && p.shield <= m.Attack {
			dgttopv := (m.Attack) - p.shield
			p.removePV(dgttopv)
			p.shield = 0
			m.PV += m.Attack * (20 / 100)
		} else {
			p.removeShield(m.Attack)
		}
	} else if CriticalChance == 10 {
		fmt.Println("Critical Hit !", m.Name, "deals to", p.nom, m.Attack*CriticalDamage, "damages")
		if p.shield == 0 {
			p.removePV(m.Attack * CriticalDamage)
			m.PV += m.Attack * (30 / 100)
		} else if p.shield > 0 && p.shield <= m.Attack*3 {
			dgttopv := (m.Attack * CriticalDamage) - p.shield
			p.removePV(dgttopv)
			p.shield = 0
			m.PV += m.Attack * (30 / 100)
		} else {
			p.removeShield(m.Attack * CriticalDamage)
		}
	} else {
		random := rand.Intn(9)
		if random == CriticalChance {
			fmt.Println("Critical Hit !", m.Name, "deals to", p.nom, m.Attack*CriticalDamage, "damages")
			if p.shield == 0 {
				p.removePV(m.Attack * CriticalDamage)
				m.PV += m.Attack * (30 / 100)
			} else if p.shield > 0 && p.shield <= m.Attack*3 {
				dgttopv := (m.Attack * CriticalDamage) - p.shield
				p.removePV(dgttopv)
				p.shield = 0
				m.PV += m.Attack * (30 / 100)
			} else {
				p.removeShield(m.Attack * CriticalDamage)
			}
		} else {
			fmt.Println(m.Name, "deals to", p.nom, m.Attack, "damages")
			if p.shield == 0 {
				p.removePV(m.Attack)
				m.PV += m.Attack * (20 / 100)
			} else if p.shield > 0 && p.shield <= m.Attack {
				dgttopv := (m.Attack) - p.shield
				p.removePV(dgttopv)
				p.shield = 0
				m.PV += m.Attack * (20 / 100)
			} else {
				p.removeShield(m.Attack)
			}
		}
	}
	fmt.Println("your remaining Shield", p.checkShield())
	fmt.Println("your remaining PVs", p.checkPV(), "/", p.pvMax)
	fmt.Println("Bosse's remaining PVs", m.PV, "/", m.PVMax)
}

// Fonction pour le pattern du monstre
func (m *monster) MonsterPattern(tour int, p *Personnage) {
	if m.Status == "Burned" {
		fmt.Println(m.Name, "is burned and takes", m.StatusEffect, "damages")
		m.PV -= m.StatusEffect
		fmt.Println(m.Name, "remaining PVs", m.PV, "/", m.PVMax)
	}
	if m.Name == "Goblin" {
		m.GoblinPattern(tour, p)
	} else if m.Name == "Orc" {
		m.OrcPattern(tour, p)
	} else if m.Name == "World Eater" {
		m.SecretBossPattern(tour, p)
	}
}
