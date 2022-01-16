package main

import (
	"fmt"
	"regexp"
	"sort"
)

type Language struct {
	Name string
}

type Rule struct {
	Literal    string
	Reg        *regexp.Regexp
	Of         *Language
	Name       string
	Translated string
	// from lowest InsertAt.Where to hightst
	ElementsAllocation []InsertAt
}

type InsertAt struct {
	Where int
	Which int
}

func (r *Rule) Translate(code string) (string, error) {
	matches := r.Reg.FindAllStringSubmatch(code, -1)
	// not the complete check for elements yet: need to see the positions of matches as well
	if len(matches[0]) < len(r.ElementsAllocation) {
		return "", fmt.Errorf("cannot find enough in code; needs %d elements; check if code follows this rule", len(r.ElementsAllocation))
	}
	// the first list of submatches is all we need
	// the first submatch is to whole declaration so we descard it
	elements := matches[0][1 : 1+len(r.ElementsAllocation)]
	translatedCode := r.Translated
	lengthOfAddedElements := 0
	for _, insertAt := range r.ElementsAllocation {
		where := insertAt.Where + lengthOfAddedElements
		translatedCode = translatedCode[:where] + elements[insertAt.Which] + translatedCode[where:]
		lengthOfAddedElements += len(elements[insertAt.Which])
	}
	return translatedCode, nil
}

func NewRule(language *Language, literal string, translated string, elementsAllocation []InsertAt, name string) (Rule, error) {
	reg, err := regexp.Compile(literal)
	if err != nil {
		return Rule{}, fmt.Errorf("can not compile literal, error: %s", err)
	}
	sort.Slice(elementsAllocation, func(i, j int) bool {
		return elementsAllocation[i].Where < elementsAllocation[j].Where
	})
	// need to do more verification
	return Rule{
		Literal:            literal,
		Reg:                reg,
		Of:                 language,
		Name:               name,
		Translated:         translated,
		ElementsAllocation: elementsAllocation,
	}, nil
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
	aLiteral := `make a ([[:alpha:]]+) variable named ([[:word:]]+) initialized to ("([^"]|(\"))*")`
	aRuleForVariable, err := NewRule(nil, aLiteral, "var   = ", []InsertAt{{Where: 4, Which: 1}, {Where: 5, Which: 0}, {Where: 8, Which: 2}}, "Sentence like explicit variable declaration")
	if err != nil {
		return "", err
	}
	translatedCode, err := aRuleForVariable.Translate(code)
	if err != nil {
		return "", err
	}
	return translatedCode, nil
}
