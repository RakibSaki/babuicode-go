package language

import (
	"babuicode/function"
	"babuicode/variable"
	"errors"
	"fmt"
	"reflect"
)

const (
	numberRegLit string = `[0-9]+(\.[0-9]*)?`
	stringRegLit string = `"[^"]*"`
	boolRegLit   string = `true|false`
	typeRegLit   string = `string|float|int|bool`
	nameRegLit   string = `(_|[a-z|A-Z])+\w*`
)

const valueRegLit string = `(` + numberRegLit + `)|(` + stringRegLit + `)|(` + boolRegLit + `)`

type Rule interface {
	Elements() []string
}

type Language struct {
	// variable declarators/definitions
	varDecs []*variable.Declaration
	// variable assignments/reassignments
	varAsns []*variable.Assignment
	// function declarators/definitions
	funcDecs []*function.Declaration
}

func New(instructions [][][][]string) (*Language, error) {
	if len(instructions) != 4 {
		return &Language{}, errors.New("instructions must be [][][][]string of length 4")
	}
	varDecs := make([]*variable.Declaration, 0)
	varAsns := make([]*variable.Assignment, 0)
	funcDecs := make([]*function.Declaration, 0)
	// validate vardec instructions and populate variable declaration rules
	for i, instruction := range instructions[0] {
		err := validateInstruction(instruction, variable.Declaration{})
		if err != nil {
			return &Language{}, errors.New(fmt.Sprintf("while validating %vth variable declaration instruction\n%s", i, err))
		}
		varDecs = append(varDecs, &variable.Declaration{Sequence: instruction[0], Keywords: instruction[1]})
	}
	// validate varAsn instructions and populate variable assignment rules
	for i, instruction := range instructions[0] {
		err := validateInstruction(instruction, variable.Assignment{})
		if err != nil {
			return &Language{}, errors.New(fmt.Sprintf("while validating %vth variable assignment instruction\n%s", i, err))
		}
		varAsns = append(varAsns, &variable.Assignment{Sequence: instruction[0], Keywords: instruction[1]})
	}
	// validate funcdec instructions and populate function declaration rules
	for i, instruction := range instructions[0] {
		err := validateInstruction(instruction, function.Declaration{})
		if err != nil {
			return &Language{}, errors.New(fmt.Sprintf("while validating %vth function declaration instruction\n%s", i, err))
		}
		funcDecs = append(funcDecs, &function.Declaration{Sequence: instruction[0], Keywords: instruction[1]})
	}
	return &Language{
		varDecs:  varDecs,
		varAsns:  varAsns,
		funcDecs: funcDecs,
	}, nil
}

// check wether instruction can create the desired rule (vardec, varasn or funcdec)
func validateInstruction(instruction [][]string, v Rule) error {
	// number of instructions (sequence, keywords, etc.)
	noInstructions := reflect.TypeOf(v).NumField()
	// check if instruction has exactly right number of instructions
	if len(instruction) != noInstructions {
		return errors.New(fmt.Sprintf("insruction must have length %v", string(noInstructions)))
	}
	// the elements that this kind of rule requires (name, type, value, etc.)
	elements := v.Elements()
	// check if sequence has the right elements
	if len(instruction[0]) != len(elements) {
		return errors.New(fmt.Sprintf("instruction must have sequence []string for %v elements", len(elements)))
	}
	// check if sequence has the right words
	wordsInSequence := make(map[string]bool)
	for _, x := range instruction[0] {
		wordsInSequence[x] = true
	}
	for _, element := range elements {
		if wordsInSequence[element] == false {
			return errors.New(fmt.Sprintf("instruction must have %s as an element in sequence", element))
		}
	}
	// check if there are exactly 4 keywords
	if len(instruction[1]) != len(elements)+1 {
		return errors.New(fmt.Sprintf("instruction must have keywords []string of length %s", len(elements)+1))
	}
	return nil
}
