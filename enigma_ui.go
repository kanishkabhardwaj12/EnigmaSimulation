package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// RunUI initializes the GUI
func RunUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Enigma Machine")
	myWindow.Resize(fyne.NewSize(500, 300))

	statusLabel := widget.NewLabel("")
	inputEntry := widget.NewEntry()
	inputEntry.SetPlaceHolder("Enter text to encrypt/decrypt...")
	outputLabel := widget.NewLabel("Output will appear here")

	// Load Enigma configuration from enigma.json
	config, err := LoadEnigmaConfig("enigma.json")
	if err != nil {
		statusLabel.SetText(fmt.Sprintf("Error: %v", err))
		log.Println(err)
	} else {
		statusLabel.SetText("Config loaded successfully")
	}

	encryptButton := widget.NewButton("Encrypt", func() {
		if inputEntry.Text == "" {
			statusLabel.SetText("Please enter a message to encrypt")
			return
		}

		encodedText := EncryptMessage(inputEntry.Text, config)
		outputLabel.SetText(fmt.Sprintf("Encrypted: %s", encodedText))
	})

	content := container.NewVBox(
		widget.NewLabel("Enigma Machine Simulator"),
		statusLabel,
		inputEntry,
		encryptButton,
		outputLabel,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
