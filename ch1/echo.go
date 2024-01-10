// print the command line args
package ch1

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	fmt.Println(strings.Join(os.Args[1:], ","))
	fmt.Println(os.Args[1:])
}
