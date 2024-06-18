package main

import (
	"flag"
	"fmt"
	"os"
	todo "todo/cmd"
)

const (
	todofile = ".todo.json"
)

func main() {
	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int("complete", 0, "mark a task as completed")
	del := flag.Int("del", 0, "delete a task")
	list := flag.Bool("list", false, "list all tasks")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todofile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

}
