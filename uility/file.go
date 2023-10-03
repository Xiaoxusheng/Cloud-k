package uility

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type File struct {
	list []string  //分片文件的集体表
	msg  chan bool //分片上传是否成功
	n    int64     //成功上传的分片个数
}

func NewFile() *File {
	return &File{
		msg: make(chan bool, 10),
		n:   0,
	}
}

// 本地分片操作
func (f *File) Burst(file *multipart.FileHeader, size int64) error {
	num := file.Size/size + 1
	f.list = make([]string, num)
	open, err := file.Open()
	if err != nil {
		return err
	}
	for i := 0; i < int(num); i++ {
		b := make([]byte, size)
		if file.Size-size*int64(i) < size {
			b = make([]byte, file.Size-size*int64(i))
		}
		rand.NewSource(time.Now().UnixNano())
		rand.Int31n(1000000)
		open.Seek(size*int64(i), 0)
		filepath := strconv.Itoa(rand.Intn(1000000)) + ".k"
		openFile, err := os.OpenFile("./file/"+filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
		f.list[i] = filepath
		if err != nil {
			log.Println(err)
			return err
		}
		open.Read(b)
		openFile.Write(b)

		openFile.Close()
	}
	open.Close()
	return nil
}

// FragmentUpload 分片上传整个过程
func (f *File) FragmentUpload(fileExt string) error {
	//初始化
	u, _ := url.Parse("https://cloud-k-1308109276.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  SECRETID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: SECRETKEY, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	id := GetUuid()
	name := id
	// 可选 opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), name, nil)
	if err != nil {
		return err
	}
	UploadID := v.UploadID
	//分片上传
	fmt.Println(UploadID)

	opt := &cos.CompleteMultipartUploadOptions{}

	// 注意，上传分块的块数最多10000块
	//开始上传
	key := "cloud-k/" + name + fileExt
	for i := 0; i < len(f.list); i++ {
		file, err := os.Open("./file/" + f.list[i])
		if err != nil {
			return err
		}
		go f.upload(client, key, UploadID, i+1, file, opt)
	}

	for {
		if int(f.n) == len(f.list) {
			break
		}
		select {
		case c := <-f.msg:
			if c {
				atomic.StoreInt64(&f.n, 1)
			} else {
				return errors.New("分片上传失败！")
			}
		}
	}
	//上传完成
	_, _, err = client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		return err
	}
	return nil
}

// 分片上传具体函数
func (f *File) upload(client *cos.Client, key string, UploadID string, index int, m *os.File, opt *cos.CompleteMultipartUploadOptions) {
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, index, bufio.NewReader(m), nil,
	)
	if err != nil {
		f.msg <- false
		log.Println(err)
	}
	PartETag := resp.Header.Get("ETag")
	//分片上传完成
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: index, ETag: PartETag},
	)
	f.msg <- true
}

//admin@cloudreve.org
//to95tQEE
