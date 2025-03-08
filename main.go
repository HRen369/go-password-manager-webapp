package main

import "fmt"

func main() {

	credentialNames := [][]string{
		{"Amazon","EasyPassowrd"},{"American Express","American5577"},
	}

	fmt.Println(credentialNames)

	// Iterate through array of items
	for i := 0; i < len(credentialNames); i++ {
		fmt.Println(i+1,"Name:",credentialNames[i][0])
		fmt.Println("Password:",credentialNames[i][1])	
	}
}