package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"example.com/main/note"
	"example.com/main/todo"
)

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("New Trick")

	//Save Note
	title, content:= getNoteData()
	note, err := note.New(title, content)
	if err != nil{
		fmt.Println(err)
	}
	outputData(note)

	//Save Todo
	todoText := getUserInput("Todo text: ")
	todo, err := todo.New(todoText)
	if err != nil{
		fmt.Println(err)
	}
	outputData(todo)
}

//Show title and content
func getNoteData() (string, string){
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

//User input
func getUserInput(prompt string) string{
	fmt.Println(prompt)
	
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	} 
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

//Interface saver
type saver interface{
	Save() error
}

//Save Data
func saveData(data saver) error{
	err := data.Save()
	if err != nil{
		fmt.Println("Saving failed")
		return err
	}
	fmt.Println("Saving succeed")
	return nil
}

//Output Inteface still be Saver Interface
type output interface{
	saver
	Display()
}

//Output Interface can be used method Saver Interface
func outputData(data output){
	data.Display()
	saveData(data)
}

func printSomething(value interface{}){
	switch value.(type){
	case int:
		fmt.Println("Integer:", value)
	case float64:
		fmt.Println("Float:", value)
	case string:
		fmt.Println("String:", value)
	}
}

func printAnything(value interface{}){
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
		return
	}
}

//Interface check value when run
func add(a, b interface{}) interface{}{
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt{
		return aInt + bInt
	}
	return 0
}

//Generics get default value when begin
func plus[T int | float64 | string] (a, b T) T{
	return a + b
}