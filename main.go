package main

import (
	"fmt"
	"strconv"
)

type Credential struct {
	siteName            string
	username            string
	passwordUnencrypted string
}

// Later will be converted to a file; Inside code for now
var testUser1 = Credential{
	siteName:            "Github",
	username:            "HTest",
	passwordUnencrypted: "ngh4rnv.e#fku",
}

var siteCredentials = []Credential{testUser1}

// Main Logic for Console App
func getSite() int {
	var option string
	fmt.Scan(&option)

	choiceInt, err := strconv.Atoi(option)

	if err != nil || choiceInt >= len(siteCredentials) {
		return -1
	}

	return choiceInt
}

func viewLogins() {
	fmt.Println("-- View Logins -- ")

	for i, userCredential := range siteCredentials {
		fmt.Printf("[%d] %s %s %s\n", i+1, userCredential.siteName, userCredential.username, userCredential.passwordUnencrypted)
	}
}

func addLogins() {
	fmt.Println("-- Add Logins --")
	fmt.Println("Site Name: ")
	fmt.Println("Username:  ")
	fmt.Println("Password:  ")

	fmt.Println("---++")
}

func deleteLogins() {
	fmt.Println("-- Delete Logins --")
	viewLogins()

	var indexToDel int = getSite()
	if indexToDel > -1 {
		siteCredentials = append(siteCredentials[:indexToDel], siteCredentials[indexToDel+1:]...)
	}

	fmt.Println("---++")
}

func consoleApp() {
	var running bool = true

	for running {
		fmt.Println("-- Go Password Manager Console App")
		fmt.Println("[1] View Login Credentials")
		fmt.Println("[2] Add Login Credentials")
		fmt.Println("[3] Delete Login Credentials")

		var option string
		fmt.Scan(&option)

		if option == "1" {
			viewLogins()
		} else if option == "2" {
			addLogins()
		} else if option == "3" {
			deleteLogins()
		} else if option == "7" {
			running = false
		}
	}

}

func main() {
	consoleApp()
}
