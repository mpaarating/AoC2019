package inputreader

import (
	"bufio"
	"os"
	"strconv"
)

func ProcessInput(path string) ([]float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var masses []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, _ := strconv.ParseFloat(scanner.Text(), 64)
		masses = append(masses, mass)
	}
	return masses, scanner.Err()
}
