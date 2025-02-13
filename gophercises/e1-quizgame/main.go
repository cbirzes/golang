package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// prompt user
	fmt.Println("Exercise 1")
	fmt.Println("Press enter to begin quiz...")
	// wait for user to hit enter to begin
	var input string
	fmt.Scanln(&input)

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

	// create global variable to track the users score
	var score int

	// run function to map/struct records
	quiz := makeQuiz(records)

	fmt.Println(quiz)

	for qnum, qa := range quiz {
		fmt.Printf("%s, %s, %s\n", qnum, qa.q, qa.a)
		// ask user question & collect response
		var response string
		fmt.Printf("%s: %s\n", qnum, qa.q)
		fmt.Scan(&response)

		if response == qa.a {
			score++
		}
	}
	// output the score
	fmt.Printf("Final Score: %d/%d", score, len(quiz))
}

func makeQuiz(records [][]string) map[string]question {
	// make a map to store each question and answer
	quiz := make(map[string]question)

	// go through each question/answer, adding them to map
	for i, record := range records {
		// create the question using the question struct we created
		question := question{q: record[0], a: record[1]}
		// add the question to the quiz using the index as the key for the question
		quiz[strconv.Itoa(i)] = question
	}
	return quiz
}

type question struct {
	q string
	a string
}

// create program to read in a quiz (problems.csv)
//  give quiz to user
//   track # of questions answered correctly
//    ask next question regarldess of correctness
//    at END, display # of correct questions & question amount
// invalid answers should be considered incorrect
// NOTE: question may have commas in it
