package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, please introduce yourself.")

	fmt.Print("Your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	fmt.Println("Are you human?")

	fmt.Print("y/n: ")
	scanner.Scan()
	human := scanner.Text()

	fmt.Println("What is your favorite programming language?")

	fmt.Print("Your answer: ")
	scanner.Scan()
	lang := scanner.Text()

	fmt.Println()
	fmt.Println("Answers:")
	fmt.Println("1. " + name)
	fmt.Println("2. " + human)
	fmt.Println("3. " + lang)
}
