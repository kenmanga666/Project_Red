package main

import (
	"fmt"
	"time"
)

// Fonction pour le loading
func (p *Personnage) loading() {
	clearConsole()
	for i := 0; i <= 100; i++ {
		fmt.Printf("Loading your character: %d%%\r", i)
		time.Sleep(40 * time.Millisecond)
		if i == 100 {
			clearConsole()
			fmt.Println("Your character is ready.")
			time.Sleep(1 * time.Second)
			clearConsole()
			p.menu()
		}
	}
}