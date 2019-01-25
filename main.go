package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type problem struct {
	question string
	answer   string
}

func main() {

	// declare program flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

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

	// counter for correct problems
	correct := 0

	// iterate over problems, asking user to answer and checking answer
	for i, v := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, v.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == v.answer {
			correct++
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
			answer:   v[1],
		}
	}

	// return result
	return result
}
