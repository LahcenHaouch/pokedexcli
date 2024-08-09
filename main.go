package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Command interface {
	CommandName() string
	CommandDescription() string
	Execute() error
}

type HelpCommand struct {
	name        string
	description string
}

func (c HelpCommand) CommandName() string {
	return c.name
}
func (c HelpCommand) CommandDescription() string {
	return c.description
}
func (c HelpCommand) Execute() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:\n")
	fmt.Println("help: Display ")

	return nil
}

type ExitCommand struct {
	name        string
	description string
}

func (c ExitCommand) CommandName() string {
	return c.name
}
func (c ExitCommand) CommandDescription() string {
	return c.description
}
func (c ExitCommand) Execute() error {
	// [todo]: implement

	return nil
}

func commandExit() error {
	return nil
}

var running bool = true
var commands map[string]Command = map[string]Command{
	"help": HelpCommand{
		name:        "help",
		description: "Displays a help message",
	},
	"exit": ExitCommand{
		name:        "exit",
		description: "Exit the Pokedex",
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for running {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		err := scanner.Err()

		if err != nil {
			log.Fatal(err)
			continue
		}

		val := scanner.Text()
		fmt.Println(val)
		// running = false
	}

}
