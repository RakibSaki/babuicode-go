package main

import (
	"fmt"
	"regexp"
)

type Language struct {
	Name string
}

func main() {
	fmt.Println("Hello, World!")
	newLanguage := Language{Name: "New Language"}
	babuiCode := `make a string variable named mood initialized to "I love this ভাপা পিঠা(Bhapa Pitha)"`
	translatedCode, err := Translate(&newLanguage, babuiCode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Translated the following code")
	fmt.Println(babuiCode)
	fmt.Printf("from %s language to the Golang code below\n", newLanguage.Name)
	fmt.Println(translatedCode)
}

func Translate(l *Language, code string) (string, error) {
	// regular expression for `make a (a possible type) variable named (a possible name) initialized to (a possible string or number or boolean)`
	aRuleForVariable, err := regexp.Compile(`make a ([[:alpha:]]+) variable named ([[:word:]]+) initialized to ("([^"]|(\"))*")`)
	if err != nil {
		return "", err
	}
	aVariableDeclaration := aRuleForVariable.FindString(code)
	matches := aRuleForVariable.FindAllStringSubmatch(aVariableDeclaration, -1)
	// the first list of submatches is all we need
	// the first submatch is to whole declaration so we descard it
	return fmt.Sprintf("var %s %s = %s", matches[0][2], matches[0][1], matches[0][3]), nil
}
