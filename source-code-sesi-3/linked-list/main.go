package main

import "fmt"

// Task represents a task in a to-do list
type Task struct {
	Title  string
	Status string
	Next   *Task
}

// ToDoList represents a to-do list
type ToDoList struct {
	Head *Task
	Tail *Task
}

func (l *ToDoList) AddTask(title, status string) {
	newTask := &Task{Title: title, Status: status, Next: nil}
	if l.Head == nil {
		l.Head = newTask
		l.Tail = newTask
	} else {
		l.Tail.Next = newTask
		l.Tail = newTask
	}
}

func (l *ToDoList) DisplayTaskAtPosition(position int) {
	current := l.Head
	count := 1
	for current != nil {
		if count == position {
			fmt.Printf("Title: %s, Status: %s\n", current.Title, current.Status)
			return
		}
		count++
		current = current.Next
	}
	fmt.Println("No task found at the specified position.")
}

func main() {
	todoList := &ToDoList{}
	todoList.AddTask("Buy groceries", "Completed")
	todoList.AddTask("Clean the house", "Incomplete")
	todoList.AddTask("Do laundry", "Incomplete")
	todoList.AddTask("Cook for dinner", "Incomplete")

	todoList.DisplayTaskAtPosition(3) // Output: Title: Do laundry, Status: Incomplete
}
