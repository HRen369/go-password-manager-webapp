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

// Helper Functions
func getSite() int {
	fmt.Printf("> ")
	var option string
	fmt.Scan(&option)

	choiceInt, err := strconv.Atoi(option)

	if err != nil || choiceInt >= len(siteCredentials)+1 || choiceInt <= 0 {
		return -1
	}

	return choiceInt - 1
}

// Console App Helper Functions
func iterateCredentialsSites() {
	for i, userCredential := range siteCredentials {
		fmt.Printf("[%d] %s\n", i+1, userCredential.siteName)
	}
}

// Console App Functions
func viewLogins() {
	fmt.Println("-- View Logins -- ")

	if len(siteCredentials) <= 0{
		fmt.Println("-- No Sites Stored --")
	}else{
		iterateCredentialsSites()

		choice := getSite()
		if choice >= 0{
			fmt.Printf("%s\n", siteCredentials[choice].siteName)
		}else{
			fmt.Println("-- Not a Valid Index --")
		}
	}
}

func addLogins() {
	fmt.Println("-- Add Logins --")

	fmt.Println("Site Name: ")
	var siteNameGiven string
	fmt.Scan(&siteNameGiven)

	fmt.Println("Username:  ")
	var usernameGiven string
	fmt.Scan(&usernameGiven)

	fmt.Println("Password:  ")
	var passwordGiven string
	fmt.Scan(&passwordGiven)

	var testUser = Credential{
		siteName:            siteNameGiven,
		username:            usernameGiven,
		passwordUnencrypted: passwordGiven,
	}

	siteCredentials = append(siteCredentials, testUser)

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
