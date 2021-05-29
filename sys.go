package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Person struct {
	name string
	age  int
}

type Command struct {
	key         string
	description string
	handler     func() error
}

func cleanTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func line() {
	for i := 0; i < 42; i++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
}

func DisplayMenu(commands []Command) {
	for {
		cleanTerminal()

		line()
		for _, command := range commands {
			fmt.Printf("%v - %v\n", command.key, command.description)
		}
		fmt.Printf("type 'exit' to end the program\n")
		line()

		var answer string
		fmt.Scanln(&answer)

		if answer == "exit" {
			cleanTerminal()
			os.Exit(0)
		}

		var commandFound bool
		for _, command := range commands {
			if command.key == answer {
				commandFound = true
				// call handler
				err := command.handler()
				if err != nil {
					fmt.Println("\n" + err.Error() + "\n")
				}
				break
			}
		}
		if commandFound == false {
			fmt.Println("Command not found.")
		}
		fmt.Print("Press 'Enter' to continue...")
		fmt.Scanln()
	}
}
