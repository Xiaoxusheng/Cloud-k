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
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func UploadFile(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDetails: err.Error(),
		})
	}
	files := form.File["file"]
	var i int64 = 0
	for _, file := range files {
		log.Println(file.Filename)
		wg.Add(1)
		go Upload(c, file, &i)
	}
	wg.Wait()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  strconv.FormatInt(i, 10) + "个文件上传成功！",
	})
}

func Upload(c *gin.Context, file *multipart.FileHeader, i *int64) {
	log.Println("进入upload")
	defer wg.Done()
	if file.Filename == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "file.Filename为空!",
		})
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
	_, err = client.Object.Put(ctx, key, bufio.NewReader(open), nil)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Warning,
			ErrorDetails: "上传失败，错误在Upload函数" + err.Error(),
		})
	}

	models.InsertFile(hash, name, ext, key, file.Size)
	atomic.AddInt64(i, 1)
}
