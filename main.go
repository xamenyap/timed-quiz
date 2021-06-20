package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var fileName string

	flag.StringVar(&fileName, "f", "problems.csv", "path to the file containing the quiz details")
	flag.Parse()

	f, err := os.Open(fileName)
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
