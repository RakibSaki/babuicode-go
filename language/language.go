package language

import (
	"babuicode/rule"
	"errors"
	"fmt"
	"regexp"
	"strconv"
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

// Compile converts babuicode to compiled code ie, code in base language. It scans for
// parts of code that could be instances of a rule and then replaces those instances
// with compiledcode as instructed by the rule.
func (l *Language) Compile(code string) (string, error) {
	finds := make([][]int, 0)
	ruleAt := make(map[int]*rule.Rule)
	// find matches to rules
	for _, rule := range l.Rules {
		for _, find := range rule.Find(code) {
			finds = append(finds, find)
			for i := find[0]; i < find[1]; i++ {
				ruleAt[i] = rule
			}
		}
	}
	lines := 1
	characters := 0
	for i := 0; i < len(code); i++ {
		characters++
		if code[i] == '\n' {
			lines++
		}
		if ruleAt[i] == nil {
			space, _ := regexp.Match(`[[:space:]]`, []byte{code[i]})
			if space {
				continue
			}
			unexpected := string(code[i])
			for j := i + 1; ruleAt[j] == nil; j++ {
				unexpected += string(code[j])
			}
			return "", UnexpectedCharacterError(lines, characters, unexpected)
		}
	}

}

func UnexpectedCharacterError(line, character int, unexpected string) error {
	return errors.New(fmt.Sprintf("Unexpected code \"%s\" at line %i, character %i", unexpected, line, character))
}

func ParseUnexpectedCharacterError(e error) (int, int, error) {
	err := e.Error()
	// verify it is an UnexpectedCharacterError
	errorPattern := `Unexpected code ".*" at line (\d+), character (\d+)`
	reg, _ := regexp.Compile(errorPattern)
	right := reg.MatchString(err)
	if !right {
		return 0, 0, errors.New("Not a valid UnexpectedCharacterError")
	}
	matches := reg.FindStringSubmatch(err)
	line, _ := strconv.Atoi(matches[1])
	character, _ := strconv.Atoi(matches[2])
	return line, character, nil
}
