package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type QA struct {
	question string
	answer   string
	index    int
}

func main() {
	records, err := readProblems("problems.csv")

	if err != nil {
		log.Fatal(err)
	}

	answers := 0

	for i, record := range records {

		ques := QA{
			question: record[0],
			answer:   record[1],
			index:    i + 1,
		}

		fmt.Printf("Question %d: %s = ", ques.index, ques.question)
		response, _ := bufio.NewReader(os.Stdin).ReadString('\n')

		if strings.TrimSuffix(response, "\n") == ques.answer {
			fmt.Printf("Correct!\n")
			answers += 1
		} else {
			fmt.Printf("Wrong!\n")
		}
	}

	fmt.Printf("You got %d/%d\n", answers, len(records))
}

func readProblems(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil

}
