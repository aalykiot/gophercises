package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func parseCSV(filename string) []Problem {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var problems []Problem

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}

		problems = append(problems, Problem{
			question: line[0],
			answer:   line[1],
		})
	}

	return problems
}

func main() {
	filename := flag.String("csv", "problems.csv", "The filename of the csv file ")
	duration := flag.Int("duration", 30, "The duration of the quiz game")
	flag.Parse()

	problems := parseCSV(*filename)

	timer := time.NewTimer(time.Duration(*duration) * time.Second)

	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %v = ", i, problem.question)

		c := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			c <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		case answer := <-c:
			if answer == problem.answer {
				correct++
			}
		}
	}

	fmt.Printf("You scored %d out of %d\n", correct, len(problems))
}
