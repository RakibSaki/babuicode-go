package function

import "babuicode/elements"

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
	return []string{elements.Name, elements.Arguments, elements.Returns, elements.Executable}
}
