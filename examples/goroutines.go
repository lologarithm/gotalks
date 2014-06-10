package main

import (
	"fmt"
	"time"
)

func answerKnockKnock() {
	fmt.Println("Who's there?")
}

func main() {
	fmt.Println("Knock Knock")
	go answerKnockKnock()
	fmt.Println("Race Condition!")
	time.Sleep(time.Second)
}
