package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"root/model/config"
	"strings"
)

//1. bug -> &&
//4. bug -> we cant add space on our values

func main() {
	//text template
	text := "This utility will walk you through creating a package.json file.\n" +
		"It only covers the most common items, and tries to guess sensible defaults.\n\n" +

		"This utility will walk you through creating a package.json file.\n" +
		"It only covers the most common items, and tries to guess sensible defaults.\n\n" +

		"See `npm help init` for definitive documentation on these fields\n" +
		"and exactly what they do.\n\n" +

		"Use `npm install <pkg>` afterwards to install a package and\n" +
		"save it as a dependency in the package.json file\n\n" +

		"Press ^C at any time to quit."

	fmt.Println(text)

	//query
	query := []struct {
		Question string
		Default  string
	}{
		{"package name:", " (express-)"},
		{"version:", " (1.0.0)"},
		{"description:", ""},
		{"entry point:", " (index.js)"},
		{"test command:", ""},
		{"git repository:", ""},
		{"keywords:", ""},
		{"author:", ""},
		{"license:", " (ISC)"},
	}

	var answer string
	for index := range query {
		answer = ""
		fmt.Print(query[index].Question, query[index].Default, " ")
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)

		if answer != "" { //set your answer
			query[index].Default = answer
		} else if query[index].Default != "" { //modify the default value
			query[index].Default = query[index].Default[2 : len(query[index].Default)-1]
		}
	}

	fileName := "package.json" //file name
	dir, err := os.Getwd()     //current directory

	check(err)

	fmt.Println("About to write to " + dir + "/" + fileName + ":\n")

	file := config.New(query[0].Default, query[1].Default, query[2].Default, query[3].Default, query[7].Default, query[8].Default)

	jp, _ := json.MarshalIndent(file, "", "   ")

	fmt.Println(string(jp) + "\n\n")

	answer = fmt.Sprintf("")
	fmt.Print("Is this OK? (yes) ")
	fmt.Scanln(&answer)

	if answer == "" || answer == "y" {
		err := ioutil.WriteFile(fileName, jp, 0644)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}
