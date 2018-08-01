package main

import (
	"strings"
	"github.com/pkg/errors"
	"flag"
	"os"
	"encoding/csv"
	"fmt"
	"bufio"
	"time"
	"math/rand"
)

const DefaultCSVPath = "problems_example.csv"

func CalculateGameResults(questionsWithAnswers [][]string, answers []string) (int, error) {
	if len(answers) != len(questionsWithAnswers){
		return 0, errors.New("answers amount is not equal to correct answers amount")
	}
	result := 0
	for index, singleQustionAndAnswer := range questionsWithAnswers {
		if strings.EqualFold(singleQustionAndAnswer[1], answers[index]){
			result++
		}
	}
	return result, nil
}

func Shuffle(input [][] string) [][]string {
	rand.Seed(time.Now().UnixNano())
	result := make([][]string, len(input))
	numbers := rand.Perm(len(input))
	for index, _ := range result {
		result[index] = input[numbers[index]]
	}
	return result
}

func ReadCSV(pathToFile string) ([][]string, error) {
	var err error
	f, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	return lines, err
}

func main(){
	var pathToCSV string
	flag.StringVar(&pathToCSV,"f", DefaultCSVPath, "input CSV file")
	var secondsOnQuiz int
	flag.IntVar(&secondsOnQuiz, "time", 30, "timer to stop quize")
	var shuffle bool
	flag.BoolVar(&shuffle, "s", false, "shuffle questions in quiz")
	flag.Parse()

	var collectedAnswers []string

	input, err := ReadCSV(pathToCSV)
	if err != nil {
		fmt.Println("Check input CSV file:", err)
		os.Exit(1)
	}
	if shuffle {
		input = Shuffle(input)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("You have ", secondsOnQuiz, " seconds to pass quiz. Press any key to start or 'q' to exit: ")
	userInput, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("Error while reading user input:", err)
		os.Exit(1)
	}
	if strings.EqualFold(strings.Replace(userInput, "\n", "", -1), "q"){
		fmt.Println("Bye")
		os.Exit(0)
	}

	go func() {
		for i := 0; i < secondsOnQuiz; i++ {
			time.Sleep(time.Second)
		}
		missingAnswersAmount := len(input) - len(collectedAnswers)
		for i := 0; i < missingAnswersAmount; i++{
			collectedAnswers = append(collectedAnswers, "")
		}
		result, _ := CalculateGameResults(input, collectedAnswers)
		fmt.Println()
		fmt.Println("Your time is up. You have scored:", result, "from", len(input), "possible")
		os.Exit(0)
	}()

	for _, question := range input {
		fmt.Print(question[0], " ")
		answer, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("Error while reading user input:", err)
			os.Exit(1)
		}
		answer = strings.Replace(answer, "\n", "", -1)
		collectedAnswers = append(collectedAnswers, answer)
	}

	result, _ := CalculateGameResults(input, collectedAnswers)
	fmt.Println("You have scored:", result, "from", len(input), "possible")
	fmt.Println("Bye")
	os.Exit(0)
}

