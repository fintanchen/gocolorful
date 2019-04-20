package main

import (
	"fmt"

	"github.com/fintanchen/gocolorful"
)

func main() {
	fmt.Println(gocolorful.Color["Green"], "hello", gocolorful.End)
	fmt.Println(gocolorful.Color["Purple"], "World", gocolorful.End)
}
