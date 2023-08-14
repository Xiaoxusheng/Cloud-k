package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestTime(tm *testing.T) {
	//现在的时间
	t := time.Now()
	//明天0点时间
	t1 := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, 0, 0, time.Local)

	times := time.NewTicker(t1.Sub(t))

	for {
		select {
		case <-times.C:
			// 记录到文件。
			//f, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02")+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			f, err := os.Create("../log/" + time.Now().Format("2006-01-02") + ".log")
			if err == nil {
				fmt.Println("创建成功！")
				//现在的时间
				t = time.Now()
				//明天0点时间
				t1 := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, 0, 0, time.Local)

				times = time.NewTicker(t1.Sub(t))
			}
			fmt.Println(err)
			f.Close()

		}
	}

}
