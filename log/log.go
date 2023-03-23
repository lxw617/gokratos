package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// 参考官方文档 https://github.com/sirupsen/logrus , https://www.cnblogs.com/jiujuan/p/15542743.html

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // 设置json里的日期输出格式
	})

	// log.SetFormatter(&log.JSONFormatter{}) // 设置为json格式

	// 设置为文本格式
	/*log.SetFormatter(&log.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		// FullTimestamp:true,
		// DisableLevelTruncation:true,
	})*/

	/*log.RegisterExitHandler(func() {
		fmt.Println("发生了fatal异常，执行一些必要的处理工作")
	})

	log.Warn("warn")
	log.Fatal("fatal")
	log.Info("info") //不会执行*/

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout) // 输入到 Stdout，默认输出到 Stderr

	//logfile, _ := os.OpenFile("./logrus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	//log.SetOutput(logfile) // 输出到文件里

	// logrus 日志一共7级别, 从高到低: panic, fatal, error, warn, info, debug, trace.
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

// go run log.go
func main() {

	// time="2023-03-22T10:13:08+08:00" level=info msg="a walrus appears" animal=walrus
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("a walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	// 设置默认字段
	requestLogger := log.WithFields(log.Fields{"request_id": "request_id", "trace_id": "trace_id"})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")

	//在 InfoLevel 级别下，定义了输出 Fatal 日志后，其后的日志都不能输出了，这是为什么？日志后面有个信息 exit status 1，
	//因为 logrus 的 Fatal 输出后，会执行 os.Exit(1)。那如果程序后面还有一些必要的程序要处理怎么办？
	//logrus 提供了 RegisterExitHandler 方法，在 fatal 异常时处理一些问题。
	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	// 共同字段输出
	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
