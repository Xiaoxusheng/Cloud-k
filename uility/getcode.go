package uility

import (
	"math/rand"
	"time"
)

func GetRandeCode() string {
	rand.Seed(time.Now().UnixNano())
	str := "1234567890"
	s := ""

	for i := 0; i < 8; i++ {
		s += string(str[rand.Intn(10)])
	}
	return s

}
