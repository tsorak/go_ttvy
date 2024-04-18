package main

import (
	"go_ttvy_input"
	"os"
)

func main() {
	ch := make(chan string)

	go go_ttvy_input.ReadToChannel(os.Stdin, ch)

	logStringChannel(ch)
}
