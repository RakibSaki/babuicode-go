package variable

import (
	"babuicode/grammar"
)

type Declaration struct {
	Sequence []string
	Keywords []string
}

type Assignment struct {
	Sequence []string
	Keywords []string
}

func (d Declaration) Elements() []string {
	return []string{grammar.Name, grammar.Type, grammar.Value}
}

func (a Assignment) Elements() []string {
	return []string{grammar.Name, grammar.Value}
}
