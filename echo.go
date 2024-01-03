// print the command line args
package main

import (
	"fmt"
	"os"
	"strings"
)

func echo() {
	fmt.Println(strings.Join(os.Args[1:], ","))
	fmt.Println(os.Args[1:])
}
