package main

import (
	"babuicode/language"
	"fmt"
)

func main() {
	code := "create a string variable called greeting initialized as \"Hi\"\nshow greeting"
	gocode, err := language.New().Compile(code)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(gocode)
	}
}
