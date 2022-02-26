package main

type Token struct {
	char  string
	color string
}

func (t *Token) string() string {
	return t.color + t.char + colorReset
}
