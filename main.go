package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {

	// declare program flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	// parse flags
	flag.Parse()

	// open csv file
	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatalf("Failed to open to CSV file: %s\n", *csvFilename)
	}

	// initialize csv reader
	reader := csv.NewReader(file)

	// read all lines in csv
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to parse provided CSV file: %s\n", err)
	}

	// convert lines into a slice of problems
	problems := parseLines(lines)

	// create a timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// counter for correct problems
	correct := 0

	// iterate over problems, asking user to answer and checking answer
	for i, v := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, v.question)
		answerC := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerC <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerC:
			if answer == v.answer {
				correct++
			}
		}
	}

	// print quiz results
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	result := make([]problem, len(lines))

	// fill in values for result
	for i, v := range lines {
		result[i] = problem{
			question: v[0],
			answer:   strings.TrimSpace(v[1]),
		}
	}

	// return result
	return result
}
