package function

import "babuicode/grammar"

type Declaration struct {
	Sequence []string
	Keywords []string
}

type Syntax struct {
	WrittenInBase bool
	Name          string
	NoOfArguments int
	NoOfReturns   int
	Sequence      []string
	Keywords      []string
}

func (d Declaration) Elements() []string {
	return []string{grammar.Name, grammar.Arguments, grammar.Returns, grammar.Executable}
}
