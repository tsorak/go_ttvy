package go_ttvy_input

import (
	"bufio"
	"fmt"
	"io"
)

func ReadToChannel(reader io.Reader, ch chan string) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		ch <- line
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading stdin:", err)
	}

	close(ch)
}
