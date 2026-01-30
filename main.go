package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


type model struct {
	msg string
}

func initializeModel() model {
	return model {
		msg: "Hello",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type) {

    
    case tea.KeyMsg:

       
        switch msg.String() {

       
        case "ctrl+c", "q":
            return m, tea.Quit

        
        
        }
    }

   
    return m, nil

}

func (m model) View() string {
	return m.msg
}


func main() {

	// fmt.Println("Welcome to my note taking app")
	p := tea.NewProgram(initializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err) 
		os.Exit(1)
	}
}