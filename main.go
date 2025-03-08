package main

import "fmt"

func main() {

	// This is mostly testing and iterating through my mistakes. I'm trying to learn Go while making this project

	credentials := [][]string{
		{"Amazon","EasyPassowrd"},{"American Express","American5577"},
	}

	// Iterate through array of items
	for i := range credentials {
		fmt.Println(i+1,"Name:",credentials[i][0])
		fmt.Println("Password:",credentials[i][1])	
	}
}