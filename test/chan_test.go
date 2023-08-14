package test

import (
	"fmt"
	"testing"
)

func TestChan(t *testing.T) {
	file := make(chan bool, 10)

	for i := 0; i < 5; i++ {
		file <- true
	}
	fmt.Println(len(file), cap(file))

	for {
		select {
		case a := <-file:
			fmt.Println(a, len(file))
		}
	}
}
