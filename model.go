package main

import "github.com/charmbracelet/bubbles/textinput"

type model struct {
	newFileInput textinput.Model
	visibleInputText bool
}