package main

import "fmt"

func doStuff(output chan string) {
	output <- "Who's There?"
}

func main() {
	fmt.Println("Knock Knock")
	output := make(chan string, 1)
	go doStuff(output)
	outputValue := <-output
	fmt.Printf("%s\n", outputValue)
	fmt.Println("Not A Race Condition!")
}
