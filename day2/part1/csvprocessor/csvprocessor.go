package csvprocessor

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ProcessCsv(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var inputs []int

	r := csv.NewReader(file)
	records, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range records {
		intVal, _ := strconv.Atoi(val)
		inputs = append(inputs, intVal)
	}

	return inputs, err
}
