package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	monitorings = 3
	delay       = 5
)

func main() {
	name := getName()
	showIntroduction(name)

	delay, monitorings := getMonitoringSettings()

	for {
		showMenu(name)
		command := readCommand()

		switch command {
		case 1:
			startMonitoring(delay, monitorings)
		case 2:
			fmt.Println("Displaying logs...")
			printLogs()
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

func getMonitoringSettings() (int, int) {
	var delay, monitorings int

	fmt.Println("Enter the delay between checks (in seconds):")
	fmt.Scan(&delay)

	fmt.Println("Enter the number of monitorings:")
	fmt.Scan(&monitorings)

	return delay, monitorings
}

func startMonitoring(delay, monitorings int) {
	fmt.Println("Monitoring...")

	sites := readSitesFromFile()
	if len(sites) == 0 {
		fmt.Println("No sites found to monitor.")
		return
	}

	for i := 0; i < monitorings; i++ {
		for _, site := range sites {
			testSite(site)
		}
		time.Sleep(time.Duration(delay) * time.Second)
		fmt.Println()
	}
}

func testSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Printf("Error accessing site: %s - %v\n", site, err)
		log.Printf("Error accessing site: %s - %v\n", site, err) // Log error
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("Site: %s was successfully loaded!\n", site)
	} else {
		fmt.Printf("Site: %s has problems. Status code: %d\n", site, resp.StatusCode)
		log.Printf("Site: %s has problems. Status code: %d\n", site, resp.StatusCode) // Log non-200 status
	}
}

func readSitesFromFile() []string {
	file, err := os.Open("sites.txt")
	if err != nil {
		log.Printf("An error occurred while opening the sites.txt file: %v\n", err)
		return nil
	}
	defer file.Close()

	var sites []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line != "" {
			sites = append(sites, line)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("An error occurred while reading the sites.txt file: %v\n", err)
			break
		}
	}
	return sites
}

func printLogs() {
	file, err := ioutil.ReadFile("log.txt")
	if err != nil {
		log.Printf("An error occurred while reading the log.txt file: %v\n", err)
		return
	}
	fmt.Println(string(file))
}
