package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/liweilicoder/LearnGolang/ch2"
)

func main() {
	//ch1.Server4()
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(1)
		}
		f := ch2.Fahrenheit(t)
		c := ch2.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, ch2.FToC(f), c, ch2.CToF(c))
	}
}
