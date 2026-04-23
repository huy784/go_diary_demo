// 日志模块
package logger

import "log"

// Info 输出信息日志
// msg: 日志消息内容
func Info(msg string) {
	log.Printf("[INFO] %s", msg)
}

// Error 输出错误日志
// msg: 日志消息内容
// err: 错误对象
func Error(msg string, err error) {
	log.Printf("[ERROR] %s: %v", msg, err)
}
