package main

import (
	"example.com/structs/notes"
)

func main() {
	// Following the SRP SOLID principle.

	var note = notes.New("Title", "Content")

	var noteFileSystem = notes.NewNoteFileSystem("../00-others")
	noteFileSystem.SaveNote(&note)

	var notePrinter = notes.NewNotePrinter()
	notePrinter.PrintNote(&note)
}

// go run app.go
