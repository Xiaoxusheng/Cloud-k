package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestChan(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int63n(100000))
	}

}
