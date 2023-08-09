package uility

import "time"

var MySigningKey = "welcome to use Cloud-k Auth:Mr.Lei time:2023/8/9 15.38"

type ErrorMessage struct {
	ErrorType        string    `json:"errorType"`        //错误类型
	ErrorDescription string    `json:"errorDescription"` //细节描述
	ErrorDetails     string    `json:"errorDetails"`     // 错误详情
	ErrorTime        time.Time `json:"errorTime"`        //时间
}

// 错误级别
const (
	Info      = 100
	Warning   = 300
	Error     = 400
	Critical  = 500
	Emergency = 999
)

var ErrorCodeToLevel = map[int]string{
	100: "Info",
	300: "Warning",
	400: "Error",
	500: "Critical",
	999: "Emergency",
}
