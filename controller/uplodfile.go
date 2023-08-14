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

var msgChan = make(chan map[string]any, 10)
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
		go Upload(file, &i)
	}
	wg.Wait()
	select {
	case msg := <-msgChan:
		fmt.Println(msg)
		c.JSON(http.StatusOK, msg)
	default:
		break
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  strconv.FormatInt(i, 10) + "个文件上传成功！",
	})

}

func Upload(file *multipart.FileHeader, i *int64) {
	log.Println("进入upload", file.Size, (float64(file.Size)/(1024*1024)) > 1)
	defer wg.Done()
	if (float64(file.Size) / (1024 * 1024)) > 50 {
		msgChan <- map[string]any{
			"code": 1,
			"msg":  "文件大小不能超过50m!",
		}
		fmt.Println(123)
		return
	}

	//if file.Filename == "" {
	//	msgChan <- map[string]any{
	//		"code": 1,
	//		"msg":  "file.Filename为空!",
	//	}
	//	return
	//}
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
		msgChan <- map[string]any{
			"code": 1,
			"msg":  "数据已经存在!",
		}
		return
	}
	models.InsertUserRepository(UserRepositorySave)
}
