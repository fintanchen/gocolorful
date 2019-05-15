// You can customize Level struct as you wish.
package main

import "github.com/fintanchen/gocolorful"

func main() {
	l := gocolorful.Level{
		Name:  "Test",
		Color: gocolorful.Magenta,
	}

	l.Print("OP")
}
