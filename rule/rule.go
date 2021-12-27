package rule

import (
	elementspkg "babuicode/elements"
	"errors"
	"fmt"
)

type Rule struct {
	Elements     []string
	BabuiCode    []string
	CompiledCode []string
}

func New(babuiCode, elements, compiledCode []string) (*Rule, error) {
	// check if arguments are []string of correct length
	if (len(babuiCode) != (len(elements)*2)+1) || (len(babuiCode) != len(compiledCode)) {
		return &Rule{}, errors.New("BabuiCode and CompiledCode both should have length one more than twice Elements' length")
	}
	for i, element := range elements {
		if !elementspkg.IsElement(element) {
			return &Rule{}, errors.New(fmt.Sprintf("%ith element %s is not a valid element", i, element))
		}
	}
	babuiElements := make(map[string]bool)
	compiledElements := make(map[string]bool)
	for i := 1; i < len(babuiCode); i++ {
		element := elements[(i-1)/2]
		if !elementspkg.Validate(babuiCode[i], element) {
			return &Rule{}, errors.New(fmt.Sprintf("%s in BabuiCode is not a valid instance of %s element", babuiCode[i], element))
		}
		babuiElements[babuiCode[i]] = true
	}
	for i := 1; i < len(compiledCode); i++ {
		if !babuiElements[compiledCode[i]] {
			return &Rule{}, errors.New(fmt.Sprintf("element instance %s in CompiledCode is not present in BabuiCode", compiledCode[i]))
		}
		delete(compiledElements, compiledCode[i])
	}
	if len(compiledElements) != 0 {
		return &Rule{}, errors.New("Not all elements in BabuiCode are present in CompiledCode")
	}
	return &Rule{
		Elements:     elements,
		BabuiCode:    babuiCode,
		CompiledCode: compiledCode,
	}, nil
}
