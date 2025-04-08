package notes

import (
	"encoding/json"
	"os"
	"strings"
)

// NoteFileWriter es responsable de guardar y cargar notas desde archivos.
type NoteFileWriter struct {
	basePath string // Directorio base para guardar las notas
}

// NewNoteFileWriter crea una nueva instancia de NoteFileWriter.
func NewNoteFileSystem(basePath string) *NoteFileWriter {
	return &NoteFileWriter{
		basePath: basePath,
	}
}

// SaveNote guarda una nota en un archivo JSON.
func (w *NoteFileWriter) SaveNote(note *Note) error {
	filename := strings.ReplaceAll(note.title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	var exportedNote *ExportedNote = note.Export()

	json, err := json.Marshal(exportedNote)
	if err != nil {
		return err
	}

	return os.WriteFile(w.basePath+"/"+filename, json, 0644)
}

type ExportedNote struct {
	Title   string
	Content string
}

func (n Note) Export() *ExportedNote {
	return &ExportedNote{
		Title:   n.title,
		Content: n.content,
	}
}
