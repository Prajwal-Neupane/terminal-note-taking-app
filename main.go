package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)




func initializeModel() model {
	return model {
		msg: "Hello",
	}
}




func main() {

	// fmt.Println("Welcome to my note taking app")
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err) 
		os.Exit(1)
	}
}