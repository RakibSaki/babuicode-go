package rule

import (
	elementspkg "babuicode/elements"
	"errors"
	"fmt"
	"regexp"
)

type Rule struct {
	Elements      []string
	BabuiCode     []string
	CompiledCode  []string
	Literal       string
	Reg           *regexp.Regexp
	HasExecutable bool
}

func New(babuiCode, elements, compiledCode []string) (*Rule, error) {
	err := CheckInstructionParams(babuiCode, elements, compiledCode)
	if err != nil {
		return &Rule{}, err
	}
	for _, element := range elements {
		if element == elementspkg.Executable {
			return &Rule{
				Elements:      elements,
				BabuiCode:     babuiCode,
				CompiledCode:  compiledCode,
				HasExecutable: true,
			}, nil
		}
	}
	literal := ""
	for i := 0; i < len(elements); i++ {
		literal += babuiCode[i*2]
		literal += elementspkg.Literals[elements[i]]
	}
	literal += babuiCode[len(elements)*2]
	reg, err := regexp.Compile(literal)
	if err != nil {
		return &Rule{}, errors.New(fmt.Sprintf("Error compiling rule literal: %s", err))
	}
	return &Rule{
		Elements:     elements,
		BabuiCode:    babuiCode,
		CompiledCode: compiledCode,
		Literal:      literal,
		Reg:          reg,
	}, nil
}

// CheckInstructionParams checks wether the given set of []string can create a new Rule
func CheckInstructionParams(babuiCode, elements, compiledCode []string) error {
	// check if []string parameters are of valid length
	if (len(babuiCode) != (len(elements)*2)+1) || (len(babuiCode) != len(compiledCode)) {
		return errors.New("BabuiCode and CompiledCode both should have length one more than twice Elements' length")
	}
	// check if all element names are valid
	for i, element := range elements {
		if !elementspkg.IsElement(element) {
			return errors.New(fmt.Sprintf("%ith element %s is not a valid element", i, element))
		}
	}
	babuiElements := make(map[string]bool)
	compiledElements := make(map[string]bool)
	// check if the elements are present in babuicode in order
	for i := 1; i < len(babuiCode); i++ {
		element := elements[(i-1)/2]
		if !elementspkg.Validate(babuiCode[i], element) {
			return errors.New(fmt.Sprintf("%s in BabuiCode is not a valid instance of %s element", babuiCode[i], element))
		}
		babuiElements[babuiCode[i]] = true
	}
	// check if exactly those elements are also present in compiledcode in any order
	for i := 1; i < len(compiledCode); i++ {
		if !babuiElements[compiledCode[i]] {
			return errors.New(fmt.Sprintf("element instance %s in CompiledCode is not present in BabuiCode", compiledCode[i]))
		}
		delete(compiledElements, compiledCode[i])
	}
	if len(compiledElements) != 0 {
		return errors.New("Not all elements in BabuiCode are present in CompiledCode")
	}
	return nil
}

// Find returns positions of all instances of this rule in babuicode. It only looks in current scope of
// code and not inside executables since only the language knows to spot executables, not its rules.
func (r *Rule) Find(code string) [][]int {
	return r.Reg.FindAllStringIndex(code, -1)
}
