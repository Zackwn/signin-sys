package main

import (
	"fmt"
	"strconv"
)

var db *Database
var commands []Command

func main() {
	DisplayMenu(commands)
}

func init() {
	db = NewDB()
	commands = []Command{
		{
			key:         "1",
			description: "Sign in a person",
			handler: func() error {
				fmt.Print("\n")
				var name, age string
				fmt.Print("Name: ")
				fmt.Scanln(&name)
				fmt.Print("Age: ")
				fmt.Scanln(&age)

				nAge, err := strconv.Atoi(age)
				if err != nil {
					return err
				}

				person := &Person{name: name, age: nAge}
				db.Insert(person)

				fmt.Print("\nSuccess!\n")
				return nil
			},
		},
		{
			key:         "2",
			description: "See all signed in persons",
			handler: func() error {
				fmt.Print("\n")
				for _, person := range db.List() {
					fmt.Printf("Name: %v \tAge: %v\n", person.name, person.age)
				}
				fmt.Print("\n")
				return nil
			},
		},
	}
}
