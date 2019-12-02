package fuel

import "math"

func CalculateFuel(moduleMass float64) (fuel int) {
	fuel = int(math.Floor(moduleMass / 3) - 2)
	return
}
