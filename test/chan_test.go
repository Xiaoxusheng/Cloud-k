package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestChan(t *testing.T) {
	rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Int63n(100000))
	}
	fmt.Println(uuid.NewV5(uuid.NewV4(), "user"))
}
