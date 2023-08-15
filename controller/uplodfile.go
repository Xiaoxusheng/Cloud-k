package controller

import (
	"Cloud-k/models"
	"Cloud-k/uility"
	"bufio"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"sync/atomic"
	"time"
)

var msgChan = make(chan int, 10)
var finShell = make(chan bool, 100)

//var wg sync.WaitGroup

func UploadFile(c *gin.Context) {
	t := time.Now()
	form, err := c.MultipartForm()
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDetails: err.Error(),
		})
	}
	files := form.File["file"]
	var i int64 = 0
	var m int64 = 0
	var n int64 = 0
	b := false
	v := false
	for _, file := range files {
		log.Println(file.Filename)
		go Upload(file, &i)
	}
	for {
		if b && v {
			break
		}
		select {
		case msg := <-msgChan:
			fmt.Println(len(msgChan))
			if msg == 1 {
				atomic.AddInt64(&m, 1)
			} else if msg == 2 {
				atomic.AddInt64(&n, 1)
			}
			//有不符合规定的文件
			if len(msgChan) == 0 && int(m+i+n) == len(files) {
				b = true
			}
		case <-finShell:
			//finShell全部完成，和int(m+i+n) == len(files)说明读取完毕
			if len(finShell) == 0 {
				v = true
			}
			if int(m+i+n) == len(files) {
				b = true
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "",
		"msg":  strconv.FormatInt(i, 10) + "个文件上传成功！" + strconv.FormatInt(m, 10) + "文件大小超过50m," + strconv.FormatInt(n, 10) + "个文件云盘中文件已存在",
	})
	t2 := time.Now()
	fmt.Println("经过时间", t2.Sub(t))

}

func Upload(file *multipart.FileHeader, i *int64) {
	log.Println("进入upload", file.Size, (float64(file.Size)/(1024*1024)) > 50)
	if (float64(file.Size) / (1024 * 1024)) > 50 {
		msgChan <- 1
		finShell <- true
		return
	}
	//判断文件是否已经存在
	open, err := file.Open()
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Warning,
			ErrorDetails: "文件打开，错误在Upload函数" + err.Error(),
		})
	}
	m := md5.New()
	if _, err := io.Copy(m, open); err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Warning,
			ErrorDetails: "ioCopy" + err.Error(),
		})
	}
	hash := fmt.Sprintf("%x", m.Sum(nil))
	ok, err := models.GetByHash(hash)
	if err != nil {
		panic(uility.ErrorMessage{
			uility.Error,
			"repository_pool表查询出错" + err.Error(),
			"GetByHash函数",
			time.Now(),
		})
	}
	if ok {
		msgChan <- 2
		finShell <- true
		return
	}
	u, _ := url.Parse("https://cloud-k-1308109276.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  uility.SECRETID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: uility.SECRETKEY, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	//文件名称
	name := uility.GetUuid()
	ext := path.Ext(file.Filename)
	key := "cloud-k/" + name + ext
	ctx := context.Background()
	log.Println("op", bufio.NewReader(open).Size())
	f, err := file.Open()
	_, err = client.Object.Put(ctx, key, bufio.NewReader(f), nil)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Warning,
			ErrorDetails: "上传失败，错误在Upload函数" + err.Error(),
		})
	}
	models.InsertFile(hash, name, ext, key, file.Size)
	atomic.AddInt64(i, 1)
	finShell <- true
	log.Println("完成")
}

func RepositorySave(c *gin.Context) {
	UserRepositorySave := new(uility.UserRepositorySave)
	err := c.BindJSON(&UserRepositorySave)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Warning,
			ErrorDetails: "解析json失败：" + err.Error(),
		})
	}
	f := models.GetByUserRepository(UserRepositorySave.UserIdentity, UserRepositorySave.RepositoryIdentity, UserRepositorySave.Name, UserRepositorySave.ParentId, UserRepositorySave.Ext)
	if f {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "数据已经存在!",
		})
		return
	}
	models.InsertUserRepository(UserRepositorySave)
}
