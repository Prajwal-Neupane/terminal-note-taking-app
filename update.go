package main

import tea "github.com/charmbracelet/bubbletea"



func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){

    var cmd tea.Cmd
	switch msg := msg.(type) {

    
    case tea.KeyMsg:

       
        switch msg.String() {

       
        case "ctrl+c", "q":
            return m, tea.Quit
        case "ctrl+n":
            m.visibleInputText = true
            return m, nil

        
        
        }
    }

    if m.visibleInputText {
        m.newFileInput, cmd = m.newFileInput.Update(msg)
    }

   
    return m, cmd

}