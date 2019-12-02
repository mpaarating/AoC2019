package main

import (
	"fmt"
	"github.com/mpaarating/aoc2019/day1/part1/fuel"
	"github.com/mpaarating/aoc2019/day1/part1/inputreader"
	"log"
)

const dataFile = "../part1/data/data.txt"

func main() {
	var totalFuel int
	modulesMass, err := inputreader.ProcessInput(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	for _, mass := range modulesMass {
		moduleFuel := fuel.CalculateFuel(mass)
		fuelMass := moduleFuel
		for fuelMass >= 0 {
			fuelMass = fuel.CalculateFuel(float64(fuelMass))
			if fuelMass > 0 {
				moduleFuel += fuelMass
			}
		}
		totalFuel += moduleFuel
	}

	fmt.Printf("Total necessary fuel is: %d \n", totalFuel)
}
