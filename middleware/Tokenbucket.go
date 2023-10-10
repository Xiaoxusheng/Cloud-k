package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
	"time"
)

type TokenBucket struct {
	capacity int64         //令牌桶容量
	rate     time.Duration //生成速率
	current  int64         //当前令牌桶令牌数量
	listTime time.Time     //上一次请求时间
}

func NewTokenBucket(cap int64, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity: cap,
		rate:     rate,
		current:  cap,
		listTime: time.Now(),
	}
}

// 取令牌
func (t *TokenBucket) tokeToken() bool {
	//当前时间
	now := time.Now()
	//	计算生成的令牌
	interval := now.Sub(t.listTime)
	//生成令牌
	GenerateToken := (interval * t.rate).Milliseconds() / 100
	//fmt.Println("生成令牌数", GenerateToken, "当前令牌总数", t.current, interval, (interval * t.rate).Milliseconds())

	// 当前总令牌
	if t.current+GenerateToken > (t.capacity) {
		t.current = t.capacity
	}
	//	1判断令牌桶是否还有令牌
	if t.current > 0 {
		//原子操作，保证线程安全
		atomic.AddInt64(&t.current, -1)
		fmt.Println("当前令牌总数", t.current)
		t.listTime = time.Now()
		return true
	}
	t.listTime = time.Now()
	return false

}

func ParseTakeToken() gin.HandlerFunc {
	tokenBucket := NewTokenBucket(50, 1)
	return func(c *gin.Context) {
		if !tokenBucket.tokeToken() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "too many requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
