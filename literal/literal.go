package literal

const (
	Number string = `[0-9]+(\.[0-9]*)?`
	String string = `"[^"]*"`
	Bool   string = `true|false`
	Type   string = `string|float|int|bool`
	Name   string = `(_|[a-z|A-Z])+\w*`
)

const Value string = `(` + Number + `)|(` + String + `)|(` + Bool + `)`
