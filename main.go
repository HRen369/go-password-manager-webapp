package main

import "fmt"

type credential struct {
	siteName            string
	username            string
	passwordUnencrypted string
}

var testUser1 = credential{
	siteName:            "Github",
	username:            "HTest",
	passwordUnencrypted: "ngh4rnv.e#fku",
}

siteCredentials := []credential{}


func viewLogins() {
	fmt.Println("-- View Logins -- ")

	for i, userCredential := range siteCredentials {
		fmt.Printf("[%d] %s %s %s\n", i, userCredential.siteName, userCredential.username, userCredential.passwordUnencrypted)
	}
}

func addLogins() {
	fmt.Println("-- Add Logins --")
	fmt.Println("Site Name: ")
	fmt.Println("Username:  ")
	fmt.Println("Password:  ")

}

func deleteLogins() {
	fmt.Println("-- Delete Logins --")
	fmt.Println(("[1] Github: HTest ngh4rnv.e#fku"))
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
		} else if option == "q" {
			running = false
		}
	}

}

func main() {
	consoleApp()
}
