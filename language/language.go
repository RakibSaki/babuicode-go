package language

import (
	"babuicode/rule"
	"errors"
	"fmt"
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
	Base  string
	Rules []*rule.Rule
}

func New(base string) Language {
	return Language{Base: base}
}

func (l *Language) AddRule(babuicode, elements, compiledcode []string) error {
	r, err := rule.New(babuicode, elements, compiledcode)
	if err != nil {
		return err
	}
	l.Rules = append(l.Rules, r)
	return nil
}

func (l *Language) Compile(code string, integers ...int) (string, error) {
	var executables int
	if len(integers) == 1 {
		executables = integers[0]
	}
	finds := make([][]int, 0)
	charactersFound := make(map[int]*rule.Rule)
	for _, rule := range l.Rules {
		for _, find := range rule.Find(code) {
			finds = append(finds, find)
			for i := find[0]; i < find[1]; i++ {
				charactersFound[i] = rule
			}
		}
	}
	lines := 1
	characters := 1
	for i := 0; i < len(code); i++ {
		if code[i] == '\n' {
			lines++
		}
		return "", errors.New(fmt.Sprintf("Unexpected character %s at line %i, character %i", code[i], lines, characters))
		characters++
	}

}
