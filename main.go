package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

// These two functions work together to handle the quiz time limit

// Start the timer, will send 'true' to specified channel once time is up
func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

// Handle running out of time
func outOfTime(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("\nTime's up!")
	os.Exit(0)
}

func main() {
	var filePath = flag.String("filepath", "problems.csv", "path to questions list")
	var timeLimit = flag.Int("timelimit", 30, "time in seconds to answer all questions")
	flag.Parse()

	// intro stuff
	fmt.Println("Welcome to the Gophercise Quiz!")
	fmt.Println("Answer as many questions correctly before time runs out!")
	fmt.Println("Questions path:", *filePath)
	fmt.Printf("Time limit: %d seconds\n", *timeLimit)

	fmt.Println("")

	fmt.Println("Press Enter when you're ready to play")
	fmt.Scanln()

	// init vars for score calculation
	var numQuestions int
	var correctQuestions int

	// start the timer
	var ch = make(chan bool)
	go timer(*timeLimit, ch)
	go outOfTime(*timeLimit, ch)

	// read questions line-by-line
	questions, err := os.Open(*filePath)
	if isError(err) {
		log.Fatal(err)
	}

	r := csv.NewReader(questions)

	// parse file line by line
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

	// calculate score and display it
	fmt.Printf("%d of %d questions answered correctly\n", correctQuestions, numQuestions)
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