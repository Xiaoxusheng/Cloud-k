package test

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestIp(t *testing.T) {
	res, err := http.Get("https://ifconfig.me/")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	all, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))

}
