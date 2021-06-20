package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var (
		fileName string
		duration time.Duration
	)

	flag.StringVar(&fileName, "f", "problems.csv", "path to the file containing the quiz details")
	flag.DurationVar(&duration, "d", 10*time.Second, "duration allowed for the quiz")
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

	// track number of correct answer
	correct := 0

	// track when the quiz is finished
	finish := make(chan struct{})

	fmt.Printf("The quiz duration is %v\n", duration)
	fmt.Println("Press any key to start the quiz")

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	go func() {
		select {
		case <-time.After(duration):
			finish <- struct{}{}
		}
	}()

	go func() {
		select {
		case <-finish:
			fmt.Printf("You answer %d/%d correctly", correct, len(records))
			os.Exit(0)
		}
	}()

	for _, rec := range records {
		fmt.Println(rec[0] + "?")
		if !sc.Scan() {
			log.Fatal("unable to read stdin")
		}

		if sc.Text() == rec[1] {
			correct++
		}
	}

	finish <- struct{}{}
}
