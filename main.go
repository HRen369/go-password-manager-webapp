package main

import "fmt"

func main() {
	var name, password string = "Amazon","1234"

	// Iterate through array of items
	for i := 0; i < 2; i++ {
		fmt.Println(i+1,"Name:",name)
		fmt.Println("Password:",password)	
	}
}