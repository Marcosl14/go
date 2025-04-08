package notes

import (
	"fmt"
)

// NotePrinter es responsable de imprimir la información de una nota a la salida estándar.
type NotePrinter struct{}

// NewNotePrinter crea una nueva instancia de NotePrinter.
func NewNotePrinter() *NotePrinter {
	return &NotePrinter{}
}

// PrintNote imprime la información de una nota.
func (p *NotePrinter) PrintNote(note *Note) {
	fmt.Println("--- Nota ---")
	fmt.Println("Título:", note.Title())
	fmt.Println("Contenido:", note.Content())
	fmt.Println("------------")
}
