package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
		scanner := bufio.NewScanner(os.Stdin)

		for {
				fmt.Print("Pokedex > ")
				if !scanner.Scan() {
						break
				}
				input := scanner.Text()
				input = strings.TrimSpace(input)
				if input == "" {
						continue
				}
				words := strings.Fields(input)
				firstWord := strings.ToLower(words[0])
				fmt.Printf("Your command was: %s\n", firstWord)
				}
}

func cleanInput(text string) []string {
		text = strings.ToLower(text)
		text = strings.TrimSpace(text)
		if text == "" {
				return []string{}
		}
		return strings.Fields(text)
}
