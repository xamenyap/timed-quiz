package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("unable to open file", err)
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal("unable to read file", err)
	}

	sc := bufio.NewScanner(os.Stdin)
	var (
		correct int
		total   int
	)

	for _, rec := range records {
		fmt.Println(rec[0] + "?")
		if !sc.Scan() {
			log.Fatal("unable to read stdin")
		}

		if sc.Text() == rec[1] {
			correct++
		}

		total++
	}

	fmt.Printf("You answer %d/%d correctly", correct, total)
}
