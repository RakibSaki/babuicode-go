package language

import (
	"babuicode/elements"
	"babuicode/rule"
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

func (l *Language) Compile(code string) (string, error) {
	executables := 0
	rules := make([]string, 0)
	for _, rule := range l.Rules {
		if rule.Literal != "" {
			
		}
		rules = append(rules, rule.Literal)
	}
	executableLiteral := elements.ExecutableLiteral(executables, rules)
}