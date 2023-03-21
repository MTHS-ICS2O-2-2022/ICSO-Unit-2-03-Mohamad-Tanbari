// Created by: Mohamad Tanbari
// Created on: Mar 2023
//
// This program displays, "Hello, World!"

package main

import "fmt"

func main() {
	// This function showws your adress
	var streetName string
	var streetNumber int

	// input
	fmt.Println("This program shows your address.")
	fmt.Println()
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street name: ")
	fmt.Scanln(&streetName)

	// output
	fmt.Println("Your address is", streetNumber, streetName, "Street.")

	// Done
	fmt.Println("\nDone.")
}
