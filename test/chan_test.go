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
	fmt.Println(int64(float64(1035993088) / (1024 * 1024)))
	for {
		select {
		case a := <-file:
			fmt.Println(a, len(file))
		}
	}
}
