package main

import "os"

type SaveI interface {
	Save() error
}

type DisplayI interface {
	Display()
}

// embeded interface
type OutputI interface {
	SaveI
	DisplayI
}

func main() {
	var note1 Note = Note{
		Title:    "Note 1",
		Text:     "This is the first note",
		Filename: "note1.txt",
	}
	saveData(note1)
	// displayData(note1) // Not allowed, because Note does not implement DisplayI

	var todo1 Todo = Todo{
		Title:    "Todo 1",
		Text:     "This is the first todo",
		Filename: "todo1.txt",
	}
	saveData(todo1)
	displayData(todo1)

	// var task1 Task = Task{
	// 	Title:    "Task 1",
	// 	Text:     "This is the first task",
	// 	Filename: "task1.txt",
	// }
	// saveData(task1) // Not allowed, because Task does not implement SaveI
}

func displayData(data OutputI) {
	data.Display()
	saveData(data)
}

func saveData(data SaveI) {
	data.Save()
}

type Note struct {
	Title    string
	Text     string
	Filename string
}

func (note Note) Save() error {
	return os.WriteFile("../00-others/"+note.Filename, []byte(note.Text), 0644)
}

type Todo struct {
	Title    string
	Text     string
	Filename string
}

func (todo Todo) Save() error {
	return os.WriteFile("../00-others/"+todo.Filename, []byte(todo.Text), 0644)
}

func (todo Todo) Display() {
	println(todo.Title)
}

type Task struct {
	Title    string
	Text     string
	Filename string
}

// go run app.go
