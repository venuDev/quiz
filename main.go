package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "A CSV file with question, answers in separted by 'coma(,)'")

	flag.Parse()
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Unable to open the file %v", *csvFileName), err)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("", err)
	}
	quizquestions := parseLines(lines)
	correct := 0
	for i, quiz := range quizquestions {
		fmt.Printf("Problem #%d :%s = ?\n", i+1, quiz.question)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == quiz.answer {
			//fmt.Printf("Correct !! %s = %s", quiz.question, quiz.answer)
			correct++
		}
	}
	fmt.Printf("You Scored %d out of %d. \n", correct, len(lines))
}

type quiz struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []quiz {
	ret := make([]quiz, len(lines))

	for i, line := range lines {
		ret[i] = quiz{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string, err error) {
	fmt.Println(msg)
	if err != nil {
		log.Fatalln(err)
	}

}
