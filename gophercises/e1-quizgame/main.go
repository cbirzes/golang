package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	// prompt user
	fmt.Println("Exercise 1")
	fmt.Println("Press enter to begin quiz. You will have 30 seconds to finish.")
	// wait for user to hit enter to begin
	var input string
	fmt.Scanln(&input)

	// start the 30 second timer
	duration := 30 * time.Second
	timer := time.NewTimer(duration)

	// Open Quiz and Read in quiz problems
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Could not read csv file.")
		return
	}

	// run function to map/struct records
	quiz := makeQuiz(records)

	// track score with score variable below
	var score int

problemloop:
	for i, question := range quiz {
		// ask user question & collect response
		fmt.Printf("%d: %s\n", i+1, question.q)
		answerChannel := make(chan string)

		go func() {
			var response string
			fmt.Scan(&response)
			answerChannel <- response
		}()

		select {
		case <-timer.C:
			fmt.Println("Timer expired!")
			break problemloop
		case response := <-answerChannel:
			if response == question.a {
				score++
			}
		}
	}
	// quiz is over, display the users final score
	fmt.Printf("Final Score: %d/%d\n", score, len(quiz))
}

func makeQuiz(records [][]string) []question {
	// make a map to store each question and answer
	quiz := make([]question, len(records))

	// go through each question/answer, adding them to map
	for i, record := range records {
		// create the question using the question struct we created
		quiz[i] = question{
			q: record[0],
			a: record[1],
		}
	}
	return quiz
}

type question struct {
	q string
	a string
}
