package ch3

import "fmt"

const GoUsage = `11111111111
1111111111111
222`

func Echo() {
	fmt.Println(GoUsage)
	fmt.Println("\xe4\xb8")

	var x rune = '\u4e16'
	fmt.Println(x)
}
