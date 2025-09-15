package cli

import (
	"fmt"
)

type Symbol string

const (
	InputSymbol   Symbol = ">"
	ResultSymbol  Symbol = "~>"
	FailureSymbol Symbol = "⚠️"
)

func Print(msg string, ident int, breakline bool, symbol Symbol) {
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

func Cursor() {
	print("> ")
}

func ReportCommandFailure(err error) {
	Print(
		fmt.Sprintf("Error: %s", err.Error()),
		0,
		true,
		FailureSymbol,
	)
}

func PreserveTyping() {
	fmt.Print("\r\033[K")
}

func Clear() {
	fmt.Print("\033[2J\033[H")
}
