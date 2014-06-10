package main

import "fmt"

// FIRST OMIT
type Traveler interface {
	Travel(string)
}

type Car struct {
	City string
}

func (c *Car) Travel(city string) {
	c.City = city
}

// ENDFIRST OMIT

func main() {
	c := &Car{City: "San Francisco"}
	fmt.Printf("City: %s\n", c.City)

	gotoBozeman(c)

	fmt.Printf("City: %s\n", c.City)
}

func gotoBozeman(t Traveler) {
	fmt.Printf("Traveling!\n")
	t.Travel("Bozeman")
}

// ENDSECOND OMIT
