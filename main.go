package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

var numQuestions int
var correctQuestions int

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func main() {
	// intro stuff
	fmt.Println("Welcome to the Gophercise Quiz!")
	fmt.Println("Answer as many questions correctly before time runs out!")

	fmt.Println("")
	fmt.Println("Getting questions from problems.csv...")

	// read questions line-by-line
	questions, err := os.Open("problems.csv")
	if isError(err) {
		log.Fatal(err)
	}

	r := csv.NewReader(questions)

	for {
		line, err := r.Read()
		// we've reached end of file, stop reading
		if err == io.EOF {
			break
		} else if err != io.EOF {
			// something wrong happens
			if isError(err) {
				return
			}
		}

		// increment one to number of questions
		numQuestions = numQuestions + 1

		question, answer := line[0], line[1]
		fmt.Printf("%s: ", question)

		// receive answer and check if it's correct
		var input string
		fmt.Scanln(&input)
		if input == answer {
			// if correct, add one to correctQuestions
			correctQuestions = correctQuestions + 1
		}
	}

	// we've finished the quiz, let's count score
	fmt.Println("Finished!")

	// calculate score and display it
	fmt.Printf("Out of %d questions, you answered %d correctly\n", numQuestions, correctQuestions)
	var score float32 = float32(correctQuestions) / float32(numQuestions) * 100
	fmt.Printf("Your total score is %.2f\n", score)

	// exit the program
	fmt.Printf("Thanks for playing!\n")
	os.Exit(0)

	// something wrong happened
	if isError(err) {
		log.Fatal(err)
	}
}
