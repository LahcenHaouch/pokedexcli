package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Command interface {
	execute() error
	getName() string
	getDescription() string
}
type BaseCommand struct {
	name        string
	description string
}

type HelpCommand struct {
	BaseCommand
}
type ExitCommand struct {
	BaseCommand
}

func (c HelpCommand) getName() string {
	return c.BaseCommand.name
}
func (c HelpCommand) getDescription() string {
	return c.BaseCommand.description
}
func (c HelpCommand) execute() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		fmt.Println(fmt.Sprintf("%s: %s", command.getName(), command.getDescription()))
	}

	fmt.Println()
	return nil
}

func (c ExitCommand) getName() string {
	return c.BaseCommand.name
}
func (c ExitCommand) getDescription() string {
	return c.BaseCommand.description
}
func (c ExitCommand) execute() error {
	running = false
	return nil
}

func commandExit() error {
	return nil
}

var running bool = true
var commands map[string]Command = map[string]Command{
	"help": HelpCommand{
		BaseCommand: BaseCommand{
			name:        "help",
			description: "Displays a help message",
		},
	},
	"exit": ExitCommand{
		BaseCommand: BaseCommand{
			name:        "exit",
			description: "Exit the Pokedex",
		},
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

		key := scanner.Text()
		command, ok := commands[key]

		if !ok {
			fmt.Println("command not found")
			continue
		}

		command.execute()
	}

}
