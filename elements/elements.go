package elements

import (
	"regexp"
)

const Value string = "value"
const Number string = "number"
const String string = "string"
const Bool string = "bool"
const Name string = "name"
const Type string = "type"
const Condition string = "condition"
const Executable string = "executable"
const Arguments string = "arguments"
const Returns string = "return values"

var Elements map[string]bool = map[string]bool{
	Value:      true,
	Name:       true,
	Type:       true,
	Condition:  true,
	Executable: true,
	Arguments:  true,
	Returns:    true,
}

var Literals map[string]string = map[string]string{
	Number:     `\d+\.\d+|\d`,
	String:     `"[^"]*"`,
	Name:       `(_|[a-z|A-Z])+\w*`,
	Type:       `string|float|int|bool`,
	Bool:       `true|false`,
	Value:      `([0-9]+(\.[0-9]*)?)|("[^"]*")|(true|false)`,
	Executable: `(\t.*\n)+`,
}

var Regexps map[string]*regexp.Regexp = map[string]*regexp.Regexp{
	Number: regexp.MustCompile(Literals[Number]),
	String: regexp.MustCompile(Literals[String]),
	Name:   regexp.MustCompile(Literals[Name]),
	Type:   regexp.MustCompile(Literals[Type]),
	Bool:   regexp.MustCompile(Literals[Bool]),
	Value:  regexp.MustCompile(Literals[Value]),
}

// IsElement function checks if a string is a valid element name.
func IsElement(s string) bool {
	return Elements[s]
}

// Validate function checks if a given string may be an instance of the specified element. It
// deosn't only check if the string contains the element but wether the whole string is the element.
func Validate(s, element string) bool {
	if !IsElement(element) {
		return false
	}
	return Regexps[element].FindString(s) == s
}
