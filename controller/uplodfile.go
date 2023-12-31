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
var finShellChan = make(chan bool, 100)

// 上传文件
func UploadFile(c *gin.Context) {
	t := time.Now()
	form, err := c.MultipartForm()
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorDescription: err.Error(),
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
	//阻塞等待所有协程完成
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
		case <-finShellChan:
			//finShell全部完成，和int(m+i+n) == len(files)说明读取完毕
			if len(finShellChan) == 0 {
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
		finShellChan <- true
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
		finShellChan <- true
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

	reader := &uility.Reader{
		Reader: f,
		Total:  file.Size,
	}

	_, err = client.Object.Put(ctx, key, reader, nil)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Warning,
			ErrorDetails: "上传失败，错误在Upload函数" + err.Error(),
		})
	}
	models.InsertFile(hash, name, ext, key, file.Size)
	atomic.AddInt64(i, 1)
	finShellChan <- true
	log.Println("完成")
}

// 同步
func RepositorySave(c *gin.Context) {
	UserRepositorySave := new(uility.UserRepositorySave)
	err := c.ShouldBindJSON(&UserRepositorySave)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Warning,
			ErrorDescription: "解析json失败：" + err.Error(),
		})
	}
	UserIdentity := c.MustGet("UserIdentity").(string)
	//查询
	f := models.GetByUserRepository(UserIdentity, UserRepositorySave.RepositoryIdentity)
	if f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "数据已经存在!",
		})
		return
	}
	UserRepositorySave.UserIdentity = UserIdentity
	models.InsertUserRepository(UserRepositorySave)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "插入成功",
	})
}

// 文件列表
func FileList(c *gin.Context) {
	//页数
	page := c.DefaultQuery("page", "1")
	//每页数量默认为20条
	number := c.DefaultQuery("number", "20")

	parent_id := c.Query("parent_id")

	if parent_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	Page, err := strconv.Atoi(page)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorTime:        time.Now(),
			ErrorDescription: err.Error(),
		})
	}
	Number, err := strconv.Atoi(number)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Error,
			ErrorTime:        time.Now(),
			ErrorDescription: err.Error(),
		})
	}
	fmt.Println(c.MustGet("UserIdentity").(string))
	fileList := models.GetFileList(Page, Number, c.MustGet("UserIdentity").(string), parent_id)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取数据成功!",
		"data": gin.H{
			"filelist": fileList,
		},
	})

}

// 更新文件名称
func UpdateFileName(c *gin.Context) {
	//
	identity := c.Query("identity")
	repository_identity := c.Query("repository_identity")
	name := c.Query("name")
	if identity == "" || name == "" || repository_identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	userIdentity := c.MustGet("UserIdentity").(string)
	k := models.GetByIdentity(identity, userIdentity, repository_identity)
	if !k {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件不存在!",
		})
		return
	}

	f := models.GetByName(name)
	if f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件名称已经存在!请更换名称",
		})
		return
	}

	models.UpDateFileName(name, identity, userIdentity)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件名称修改成功!",
	})

}

// 创造文件夹
func CreateFolder(c *gin.Context) {
	parent_id := c.Query("parent_id")
	name := c.Query("name")
	if parent_id == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	userIdentity := c.MustGet("UserIdentity").(string)
	ok := models.GetByNameParentId(name, parent_id)
	if ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "该名称已经存在!请更换名称",
		})
		return
	}
	Parent_id, err := strconv.Atoi(parent_id)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        err.Error(),
			ErrorDescription: "user_basic表插入出错" + err.Error(),
			ErrorTime:        time.Now(),
		})
	}
	models.InsertFolder(userIdentity, name, Parent_id)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件夹创建成功!",
	})
}

// 删除文件
func DeleteFile(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	userIdentity := c.MustGet("UserIdentity").(string)
	//查询文件是否存在
	f, _ := models.GetByIdentityUserIdentity(identity, userIdentity)
	if !f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件不存在!",
		})
		return
	}

	models.DeleteFile(identity, userIdentity)
	c.JSON(http.StatusBadRequest, gin.H{
		"code": 200,
		"msg":  "删除成功!",
	})
}

// MoveFile 移动文件
func MoveFile(c *gin.Context) {
	//目的文件夹的唯一id
	parent_identity := c.Query("parent_identity")
	//文件的唯一id
	identity := c.Query("identity")

	if parent_identity == "" || identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	//查询目的文件夹是否存在
	userIdentity := c.MustGet("UserIdentity").(string)
	f, user := models.GetByIdentityUserIdentity(parent_identity, userIdentity)
	if !f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件夹不存在!",
		})
		return
	}
	ok, _ := models.GetByIdentityUserIdentity(identity, userIdentity)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件不存在!",
		})
		return
	}

	models.UpdateFileParentId(identity, user.Id)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "文件移动成功!",
	})

}

// 文件夹列表
func FolderList(c *gin.Context) {
	userIdentity := c.MustGet("UserIdentity").(string)
	data := models.GetByUseIdentityRepositoryIdentityList(userIdentity, "")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "获取文件夹列表成功!",
		"data": gin.H{
			"folder_list": data,
		},
	})
}

// 修改文件夹名称
func UpdateFolder(c *gin.Context) {
	userIdentity := c.MustGet("UserIdentity").(string)
	name := c.Query("name")
	identity := c.Query("identity")
	if identity == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	//identity是否存在
	//查询文件是否存在
	f, file := models.GetByIdentityUserIdentity(identity, userIdentity)
	if !f {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件不存在!",
		})
		return
	}
	if file.RepositoryIdentity == "" {
		models.UpdateFlor(identity, name)
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "修改成功！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "不是文件夹！",
		})
	}

}

// 下载文件，目前仅支持单线程下载
func DownloadFile(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
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
	//判断是否存在
	ok, file := models.GetByRepositoryPool(identity)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "文件不存在!",
		})
		return
	}
	// 1.通过响应体获取对象
	name := "cloud-k/" + file.Name + file.Ext
	resp, err := client.Object.Get(context.Background(), name, nil)
	if err != nil {
		fmt.Println(err)
	}
	g, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	//写入响应体
	_, err = c.Writer.Write(g)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:    uility.Error,
			ErrorDetails: "写入出错" + err.Error(),
			ErrorTime:    time.Now(),
		})
	}
	//设置响应头状态码
	c.Writer.WriteHeader(http.StatusOK)
}

// 初始化分片传
func UploadPart(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	client := uility.GetClient()
	if client == nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "获取client失败",
			ErrorTime:        time.Now(),
		})
	}
	key := "cloud-k/" + uility.GetUuid() + path.Ext(filename)
	// 可选 opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		panic(err)
	}
	UploadID := v.UploadID

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "分片上传初始化成功！",
		"data": gin.H{
			"UploadId": UploadID,
			"key":      key,
		},
	})

	fmt.Println(UploadID)
}

// 分片上传
func FragmentUpload(c *gin.Context) {
	key := c.Query("key")
	UploadID := c.Query("UploadID")
	partNumber := c.Query("partNumber")
	file, err := c.FormFile("file")

	if key == "" || UploadID == "" || partNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "必填参数不能为空!",
		})
		return
	}
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "获取表单文件失败",
			ErrorTime:        time.Now(),
		})
	}
	PartNumbers, err := strconv.Atoi(partNumber)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "转换失败！",
			ErrorTime:        time.Now(),
		})
	}
	open, err := file.Open()
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "打开表单文件失败",
			ErrorTime:        time.Now(),
		})
	}

	client := uility.GetClient()
	if client == nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "获取client失败",
			ErrorTime:        time.Now(),
		})
	}
	// opt 可选
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, PartNumbers, bufio.NewReader(open), &cos.ObjectUploadPartOptions{
			ContentLength: file.Size,
		},
	)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "分片上传失败" + err.Error(),
			ErrorTime:        time.Now(),
		})
	}
	PartETag := resp.Header.Get("ETag")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "分片上传成功！",
		"data": gin.H{
			"PartETag":   PartETag,
			"PartNumber": PartNumbers,
		},
	})

	fmt.Println(PartETag)

}

// 分片上传完成
func UploadsCompletion(c *gin.Context) {
	key := c.Query("key")
	UploadID := c.Query("UploadID")
	P := c.QueryMap("p")

	client := uility.GetClient()
	if client == nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "获取client失败",
			ErrorTime:        time.Now(),
		})
	}
	opt := &cos.CompleteMultipartUploadOptions{}

	for k, v := range P {
		atoi, err := strconv.Atoi(k)
		if err != nil {
			return
		}
		opt.Parts = append(opt.Parts, cos.Object{
			PartNumber: atoi, ETag: v},
		)
	}

	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		panic(uility.ErrorMessage{
			ErrorType:        uility.Emergency,
			ErrorDescription: "分片上传完成失败",
			ErrorTime:        time.Now(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "分片上传完成！",
	})

}

//1 816e8d3b2e8f234d47be0c098cb96dce
//2 e6acb3d6a70a21c40d993a3a5a5acc22
