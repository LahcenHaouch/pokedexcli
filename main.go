package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Command interface {
	execute() error
	getBaseCmd() BaseCommand
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

func (c HelpCommand) execute() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range commands {
		baseCmd := command.getBaseCmd()
		fmt.Println(fmt.Sprintf("%s: %s", baseCmd.name, baseCmd.description))
	}

	fmt.Println()
	return nil
}
func (c HelpCommand) getBaseCmd() BaseCommand {
	return c.BaseCommand
}
func (c ExitCommand) execute() error {
	running = false
	return nil
}
func (c ExitCommand) getBaseCmd() BaseCommand {
	return c.BaseCommand
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
