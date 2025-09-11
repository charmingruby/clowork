package command

import (
	"fmt"
)

type Symbol string

const (
	InputSymbol         Symbol = ">"
	ResultSymbol        Symbol = "~>"
	ListTopSymbol       Symbol = ",-===-~"
	ListSeparatorSymbol Symbol = "|-===-~"
	ListBottomSymbol    Symbol = "'-===-~"
	FailureSymbol       Symbol = "⚠️"
)

func print(msg string, ident int, breakline bool, symbol Symbol) {
	var identation string

	if ident == 0 {
		identation = string(symbol)
	} else {
		for i := range ident {
			identation += "  "

			if i == ident-1 {
				identation += string(symbol)
				continue
			}
		}
	}

	if breakline {
		fmt.Printf("%s %s \n", identation, msg)
		return
	}

	fmt.Printf("%s %s", identation, msg)
}

func ReportFailure(err error) {
	print(
		fmt.Sprintf("Error: %s", err.Error()),
		0,
		true,
		FailureSymbol,
	)
}

func list(fn func()) {
	print("", 0, true, ListTopSymbol)
	fn()
	print("", 0, true, ListBottomSymbol)
}

func maybeSeparate(currIdx, size int) {
	if currIdx != size-1 {
		print("", 0, true, ListSeparatorSymbol)
	}
}
