package test

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	sprintf := fmt.Sprintf("%x", md5.Sum([]byte("dnjkfndsnof")))
	fmt.Println(sprintf)
}
