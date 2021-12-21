package main

import (
	"babui/language"
	"fmt"
)

// func (l *LangDef) ScanKeywords() {
// 	keywords := make(map[string]bool)
// 	for _, def := range l.varDefs {
// 		for _, keyword := range def.keywords {
// 			if !keywords[keyword] {
// 				keywords[keyword] = true
// 			}
// 		}
// 	}
// 	l.keywords = make([]string, 0)
// 	for keyword := range keywords {
// 		l.keywords = append(l.keywords, keyword)
// 	}
// }

func main() {
	fmt.Println("Hello World")
	lang, err := language.New(true, language.VarDef{Order: []int{1, 2, 3}, Keywords: [4]string{"", "will be a", "like", ""}})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(lang.Compile("food will be a string like \"rice\""))
	}
}
