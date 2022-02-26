package main

type Token struct {
	char  string
	color string
}

func (t Token) String() string {
	return t.color + t.char + colorReset
}
