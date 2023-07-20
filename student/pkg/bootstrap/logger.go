package bootstrap

import (
	"context"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/spf13/viper"
	"os"
	"student/pkg/bootstrap/logx"
	"student/pkg/constant"
	"student/pkg/util"
	"time"
)

func NewLogger() (logger klog.Logger, cleanup func()) {
	mode := viper.GetString("mode")
	switch mode {
	case constant.MODE_DEFAULT:
		logger, cleanup = logx.NewZap()
	case constant.MODE_lOCAL:
		logger, cleanup = logx.NewLogrus()
	default:
		return nil, nil
	}

	// 日志初始化
	localIP, _ := util.GetLocalIP()
	id, _ := os.Hostname()
	logger = klog.With(logger,
		"@source", localIP,
		"@traceId", tracing.TraceID(),
		"@spanId", tracing.SpanID(),
		"@timestamp", _timestamp(),
		"@date", klog.Timestamp("2006-01-02 15:04:05"),
		"@caller", klog.DefaultCaller,
		"@serviceId", id,
	)
	return
}

func _timestamp() klog.Valuer {
	return func(ctx context.Context) interface{} {
		return time.Now().Unix()
	}
}
