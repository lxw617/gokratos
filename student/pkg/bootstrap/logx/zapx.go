package logx

import (
	kzap "github.com/go-kratos/kratos/contrib/log/zap/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"log"
)

func NewZap() (klog.Logger, func()) {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()
	return kzap.NewLogger(logger), func() { log.Printf("\u001B[31m%s\u001B[0m\n", "logrus logger graceful close") }
}
