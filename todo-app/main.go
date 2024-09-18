package main

import "fmt"

var tasks []Task

type Task struct {
	Name     string
	Complete bool
}

func AddTask(name string) {
	task := Task{Name: name, Complete: false}
	tasks = append(tasks, task)
	fmt.Println("Added:", name)
}

func ListTask() {
	fmt.Println("Task:")
	for i, task := range tasks {
		status := "Incomplete"
		if task.Complete {
			status = "Complete"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Name, status)
	}
}

func CompleteTask(index int) {
	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks[index-1].Complete = true
	fmt.Println("Marked task as complete:", tasks[index-1].Name)
}

func DeleteTask(index int) {
	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Deleted task", index)
}

func main() {
	var choice int
	fmt.Println("To-Do App")
	for {
		fmt.Println("\nTo-Do List:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Task")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose and option: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Enter task name: ")
			fmt.Scan(&name)
			AddTask(name)
		case 2:
			ListTask()
		case 3:
			var index int
			fmt.Print("Enter task number to complete: ")
			fmt.Scan(&index)
			CompleteTask(index)
		case 4:
			var index int
			fmt.Print("Enter task number to delete: ")
			fmt.Scan(&index)
			DeleteTask(index)

		case 5:
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
