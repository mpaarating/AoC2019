package main

import (
	"fmt"
	"github.com/mpaarating/aoc2019/day1/part1/fuel"
	"github.com/mpaarating/aoc2019/day1/part1/inputreader"
	"log"
)

const dataFile = "data/data.txt"

func main() {
	var totalFuel int
	modulesMass, err := inputreader.ProcessInput(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, mass := range modulesMass {
		totalFuel += fuel.CalculateFuel(mass)
	}

	fmt.Printf("Total necessary fuel is: %d \n", totalFuel)
}
