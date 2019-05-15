// This package provide usage of default function.
package main

import "github.com/fintanchen/gocolorful"

func main() {
	gocolorful.Info("Info... This message is Blue.")
	gocolorful.Warning("Warning... This message is Yellow.")
	gocolorful.Debug("Debug... This message is White.")
	gocolorful.Err("Err... This message is Red")
}
