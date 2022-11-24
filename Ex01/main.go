package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	csvName = "problems.csv"
	timeout = 2
)

func main() {
	lines := preprocessing()
	problems := parseLines(lines)

	timeLimit := flag.Int("limit", timeout, "the time limit of quiz in secs")
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = ", i+1, problem.q)
		ansChan := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			ansChan <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d. \n", correct, len(problems))
			return
		case ans := <-ansChan:

			if ans == problem.a {
				correct++
			}
		}
	}
}

func preprocessing() [][]string {
	csvFileName := flag.String("csv", csvName, "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Filed to open the csv file %s", *csvFileName))
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	return lines
}

func parseLines(lines [][]string) []problem {
	prob := make([]problem, len(lines))

	for i, line := range lines {
		prob[i] = problem{line[0], strings.TrimSpace(line[1])}
	}

	return prob
}

type problem struct {
	q, a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
