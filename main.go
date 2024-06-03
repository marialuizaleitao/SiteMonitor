package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := getName()
	showIntroduction(name)

	for {
		showMenu(name)
		command := readCommand()

		switch command {
		case 1:
			// startMonitoring()
		case 2:
			fmt.Println("Displaying logs...")
			// printLogs()
		case 0:
			fmt.Println("Exiting the program")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
			os.Exit(-1)
		}
	}
}

func getName() string {
	fmt.Println("Please enter your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	return strings.TrimSpace(name)
}

func showIntroduction(name string) {
	version := 1.1
	fmt.Printf("Hello, %s\n", name)
	fmt.Printf("This program is version %.1f\n", version)
}

func showMenu(name string) {
	fmt.Printf("Welcome, %s\n", name)
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Display Logs")
	fmt.Println("0- Exit Program")
}

func readCommand() int {
	var commandRead int
	fmt.Scan(&commandRead)
	fmt.Printf("The chosen command was %d\n", commandRead)
	return commandRead
}
