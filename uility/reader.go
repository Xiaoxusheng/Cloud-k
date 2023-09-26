package uility

import (
	"fmt"
	"io"
)

type Reader struct {
	io.Reader
	Current int64
	Total   int64
}

func (r *Reader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	if err != nil {
		return 0, err
	}
	r.Current += int64(n)
	fmt.Printf("\r进度 %.2f%%", float64(r.Current*10000/r.Total)/100)
	return
}
