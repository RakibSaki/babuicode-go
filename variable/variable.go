package variable

import (
	"babuicode/elements"
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
	return []string{elements.Name, elements.Type, elements.Value}
}

func (a Assignment) Elements() []string {
	return []string{elements.Name, elements.Value}
}
