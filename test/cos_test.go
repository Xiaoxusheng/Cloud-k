package test

import (
	"Cloud-k/uility"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"testing"
)

func TestUpload(t *testing.T) {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由 bucketname-appid 组成，appid 必须填入，可以在 COS 控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://cloud-k-1308109276.cos.ap-nanjing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  uility.SECRETID,  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
			SecretKey: uility.SECRETKEY, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考 https://cloud.tencent.com/document/product/598/37140
		},
	})
	//name := "cloud-k/1.jpg"
	//// 2.通过本地文件上传对象
	//_, err := c.Object.PutFromFile(context.Background(), name, "1.jpg", nil)
	//if err != nil {
	//	panic(err)
	//}
	f := "1.jpg"

	name := "cloud-k/" + f
	ctx := context.Background()

	// 1. 通过普通方式上传对象
	_, err, _ := c.Object.Upload(ctx, name, f, nil)
	if err != nil {
		panic(err)
	}

}
