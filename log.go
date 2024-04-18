package main

import (
	"fmt"
)

func logStringChannel(ch <-chan string) {
	for {
		value, ok := <-ch

		if !ok {
			fmt.Println("Channel has been closed")
			break
		}

		fmt.Println(value)
	}
}
