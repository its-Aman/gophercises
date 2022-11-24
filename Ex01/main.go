package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := preprocessing()
	problems := parseLines(lines)
	startProblemSet(problems)
}

func preprocessing() [][]string {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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

func startProblemSet(problems []problem) {
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, problem.q)

		var ans string
		fmt.Scanf("%s\n", &ans)

		if ans == problem.a {
			// fmt.Println("Correct !")
			correct++
			// } else {

		}
	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
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
