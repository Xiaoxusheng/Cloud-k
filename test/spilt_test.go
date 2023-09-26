package test

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestSpilt(t *testing.T) {
	var size int64 = 1024 * 1024 * 50
	stat, err := os.Stat("./2.mp4")
	if err != nil {
		return
	}
	num := stat.Size()/size + 1
	file, err := os.OpenFile("./2.mp4", os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < int(num); i++ {
		b := make([]byte, size)
		if stat.Size()-size*int64(i) < size {
			b = make([]byte, stat.Size()-size*int64(i))
		}
		file.Seek(size*int64(i), 0)
		openFile, err := os.OpenFile("./"+strconv.Itoa(i)+".k", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
		if err != nil {
			log.Println(err)
			return
		}
		file.Read(b)
		openFile.Write(b)

		openFile.Close()
	}
	file.Close()

	fmt.Println(stat.Size())
	merge()
}

func merge() {
	var size int64 = 1024 * 1024 * 50
	stat, err := os.Stat("./2.mp4")
	if err != nil {
		return
	}
	num := stat.Size()/size + 1
	files, err := os.OpenFile("./text.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < int(num); i++ {
		file, err := os.OpenFile("./"+strconv.Itoa(i)+".k", os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Println(file)
			return
		}
		fmt.Println("./" + strconv.Itoa(i) + ".k")
		//all, err := io.ReadAll(file)
		//if err != nil {
		//	return
		//}
		//files.Write(all)
		io.Copy(files, files)
		file.Close()
	}
	files.Close()
}
