package gocolorful

import (
	"fmt"
	"log"
)

// Level :
// a level should have two property.
// 		Name and Color
// Name should be a valid Name
// Color see types.go
// You also can customize your Level
type Level struct {
	Name  string
	Color Color
}

// Print is a low-level function for print message to terminal
// you can create your Level structure and give them name and color
func (l Level) Print(v ...interface{}) {

	bandage := "[" + l.Name + "]"

	if l.Name == "ERR" {
		log.Fatalln(l.Color, bandage, fmt.Sprint(v...), End)
	}

	log.Println(l.Color, bandage, fmt.Sprint(v...), End)
}
