package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// struct to hold a question
type question struct {
	text   string `Text of the question, e.g. "1+1" or "Capital of Japan"`
	answer string `Answer of the question, e.g. "2" or "Tokyo"`
}

// init vars for score calculation
var numQuestions int
var correctQuestions int

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
	showResults()
	fmt.Println("Better luck next time!")
	os.Exit(0)
}

// Make a string easier to compare by removing all spaces and converting to lowercase
func prepString(s string) string {
	var result = s

	result = strings.ToLower(result)
	result = strings.ReplaceAll(result, " ", "")

	return result
}

func showResults() {
	fmt.Printf("%d of %d questions answered correctly\n", correctQuestions, numQuestions)
	var score float32 = float32(correctQuestions) / float32(numQuestions) * 100
	fmt.Printf("Your total score is %.2f\n", score)
}

// get questions
func getQuestions(path string) []question {
	var questions []question

	// try open file
	file, err := os.Open(path)
	if isError(err) {
		log.Fatal(err)
	}

	// read file line by-line
	reader := csv.NewReader(file)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != io.EOF {
			if isError(err) {
				log.Fatal(err)
			}
		}

		// parse each line and setup questions for the quiz
		var text, answer = line[0], prepString(line[1])
		var question = question{text, answer}
		questions = append(questions, question)
	}

	return questions
}

func main() {
	var filePath = flag.String("file", "problems.csv", "path to questions list")
	var timeLimit = flag.Int("time", 30, "time in seconds to answer all questions")
	flag.Parse()

	// intro stuff
	fmt.Println("Welcome to the Gophercise Quiz!")
	fmt.Println("Answer as many questions correctly before time runs out!")
	fmt.Println("Questions path:", *filePath)
	fmt.Printf("Time limit: %d seconds\n", *timeLimit)

	fmt.Println("")

	fmt.Println("Press Enter when you're ready to play")
	fmt.Scanln()

	// start the timer
	var ch = make(chan bool)
	go timer(*timeLimit, ch)
	go outOfTime(*timeLimit, ch)

	// read questions line-by-line
	var questions = getQuestions(*filePath)
	numQuestions = len(questions)

	// Loop through all questions
	scanner := bufio.NewScanner(os.Stdin)
	for _, question := range questions {
		// display question
		fmt.Printf("%s: ", question.text)

		// wait for user input
		if scanner.Scan() {
			var input string = scanner.Text()
			input = prepString(input)
			if input == question.answer {
				correctQuestions = correctQuestions + 1
			}
		}
	}

	// calculate score and display it
	showResults()

	// exit the program
	fmt.Printf("Thanks for playing!\n")
	os.Exit(0)
}
