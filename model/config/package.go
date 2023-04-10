package config

import "fmt"

type Testing struct {
	Test string `json:"test"`
}

type Package struct {
	Name        string  `json:"name"`
	Version     string  `json:"version"`
	Description string  `json:"description"`
	Main        string  `json:"main"`
	Test        Testing `json:"scripts"`
	Author      string  `json:"author"`
	License     string  `json:"license"`
}

func New(Name, Version, Description, Main, Author, License string) *Package {
	return &Package{
		Name:        Name,
		Version:     Version,
		Description: Description,
		Main:        Main,
		Test:        newTest(),
		Author:      Author,
		License:     License,
	}
}

func newTest() Testing {
	return Testing{
		Test: fmt.Sprintf("echo \"Error: no test specified\" && exit 1"),
	}
}
