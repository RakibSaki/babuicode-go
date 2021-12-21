package language

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	numberRegLit string = `[0-9]+(\.[0-9]*)?`
	stringRegLit string = `"[^"]*"`
	boolRegLit   string = `true|false`
	typeRegLit   string = `string|float|int|bool`
	nameRegLit   string = `(_|[a-z|A-Z])+\w*`
)

const valueRegLit string = `(` + numberRegLit + `)|(` + stringRegLit + `)|(` + boolRegLit + `)`

type VarDef struct {
	// order of name, type and value in declaration
	// []{2, 1, 3} means first is type, then name and lastly value
	Order    []int
	Keywords [4]string
}

func (v *VarDef) Reg(l *Language) (*regexp.Regexp, error) {
	words := make([]string, 0)
	if v.Keywords[0] != "" {
		words = append(words, v.Keywords[0])
	}
	for i := 0; i < 3; i++ {
		switch v.Order[i] {
		case 1:
			words = append(words, nameRegLit)
		case 2:
			words = append(words, typeRegLit)
		case 3:
			words = append(words, valueRegLit)
		}
		if v.Keywords[i+1] != "" {
			words = append(words, v.Keywords[i+1])
		}
	}
	literal := `(` + strings.Join(words, `) (`) + `)`
	fmt.Println("regular expression literal is")
	fmt.Println(literal)
	return regexp.Compile(literal)
}

func (v *VarDef) Compile(code string) string {
	var ntv []string = []string{code}
	for i := 0; i < len(v.Keywords); i++ {
		if v.Keywords[i] == "" {
			continue
		}
		fmt.Println("ntv is")
		fmt.Println(ntv)
		fmt.Println("keyword is", v.Keywords[i])
		var temp []string = []string{}
		for j := 0; j < len(ntv); j++ {
			temp = append(temp, strings.Split(ntv[j], v.Keywords[i])...)
		}
		ntv = temp
	}
	fmt.Println("ntv is")
	fmt.Println(ntv)
	var name, typ, value string
	for i := 0; i < 3; i++ {
		trimmed := strings.TrimSpace(ntv[i])
		switch v.Order[i] {
		case 1:
			name = trimmed
		case 2:
			typ = trimmed
		case 3:
			value = trimmed
		}
	}
	return "var " + name + " " + typ + " = " + value
}

// Essentially the whole new language. Contains the set
// of instructions dictating how the language should work
type Language struct {
	Typed    bool
	VarDef   VarDef
	Keywords []string
	// ifDefs []ifDef
	// numForDefs []numForDef
	// arrForDefs []arrForDef
	// dictForDefs []dictForDef
	// funcDefs []funcDef
	// moreDefs []moreDef
}

func New(typed bool, v VarDef) (Language, error) {
	result := Language{Typed: typed, VarDef: v}
	for _, keyword := range result.VarDef.Keywords {
		result.Keywords = append(result.Keywords, keyword)
	}
	return result, nil
}

func (l *Language) Compile(code string) (string, error) {
	varReg, err := l.VarDef.Reg(l)
	if err != nil {
		return "", err
	}
	fmt.Println("code is")
	fmt.Println(code)
	declaration := string(varReg.Find([]byte(code)))
	fmt.Println("declaration is")
	fmt.Println(declaration)
	compiled := l.VarDef.Compile(declaration)
	return compiled, nil
}
