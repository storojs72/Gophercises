package main

import (
	"net/http"
	"fmt"
	"github.com/gophercises/3_choose_your_own_adventure/dependencies"
	"io/ioutil"
	"encoding/json"
	"bufio"
	"os"
	"strconv"
	"flag"
)

func main(){
	consoleAppFlag := flag.Bool("c", false, "use TRUE to start console application")
	flag.Parse()

	if *consoleAppFlag {
		ConsoleApp()
	} else {
		WebApp()
	}
}
func ConsoleApp() {
	gopherObject := &dependencies.Gopher{}

	bytes, err := ioutil.ReadFile("3_choose_your_own_adventure/dependencies/gopher.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, gopherObject)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewScanner(os.Stdin)


	meta := "intro"
	for {
		options := PrintStruct(gopherObject, meta)
		if options == nil {
			fmt.Println("Bye")
			os.Exit(0)
		}

		for {
			reader.Scan()
			userChoice, err := strconv.Atoi(reader.Text())
			if err != nil || len(options) <= userChoice - 1 {
				fmt.Println("impossible variant")
			} else {
				meta = options[userChoice - 1].Arc
				break
			}
		}


	}

}

func PrintStruct(value interface {}, meta string) []struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
} {
	convertedValue, ok := value.(*dependencies.Gopher)
	if !ok {
		panic(value)
	}

	switch meta {
	case "intro":
		Print(convertedValue.Intro)
		return convertedValue.Intro.Options
	case "new-york":
		Print(convertedValue.NewYork)
		return convertedValue.NewYork.Options
	case "denver":
		Print(convertedValue.Denver)
		return convertedValue.Denver.Options
	case "debate":
		Print(convertedValue.Debate)
		return convertedValue.Debate.Options
	case "sean-kelly":
		Print(convertedValue.SeanKelly)
		return convertedValue.SeanKelly.Options
	case "mark-bates":
		Print(convertedValue.MarkBates)
		return convertedValue.MarkBates.Options
	case "home":
		fmt.Println("************************************************")
		fmt.Println(convertedValue.Home.Title)
		fmt.Println()
		for _, storyPart := range convertedValue.Home.Story{
			fmt.Println(storyPart)
		}
		fmt.Println()
		return nil
	default:
		panic("Unexpected")
		return nil
	}
}
func Print(value interface {}) {
	convertedValue, ok := value.(struct {
		Title   string   `json:"title"`
		Story   []string `json:"story"`
		Options []struct {
			Text string `json:"text"`
			Arc  string `json:"arc"`
		} `json:"options"`
	})

	if !ok {
		panic(value)
	}

	fmt.Println("************************************************")
	fmt.Println(convertedValue.Title)
	fmt.Println()

	for _, storyPart := range convertedValue.Story{
		fmt.Println(storyPart)
	}

	fmt.Println()
	for index, option := range convertedValue.Options{
		fmt.Println(option.Text, "To continue, press", index + 1)
	}
}



func WebApp(){
	handler := dependencies.NewMainHandler()
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}
