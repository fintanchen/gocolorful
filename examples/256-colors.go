// This package will print all 256 colors in terminal
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for row := 0; row < 16; row++ {
		for col := 0; col < 16; col++ {
			code := row*16 + col
			fmt.Fprint(os.Stdout, "\u001b[38;5;"+strconv.Itoa(code)+"m", code)
			fmt.Fprint(os.Stdout, "\u001b[0m\t")
		}
		fmt.Print("\n")
	}
}
