package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)




func initializeModel() model {

	// initialize new input file
	ti := textinput.New()
	ti.Placeholder = "Name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model {
		newFileInput: ti,
		visibleInputText: false,
		
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