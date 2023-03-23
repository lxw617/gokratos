package logx

import (
	klogrus "github.com/go-kratos/kratos/contrib/log/logrus/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

func NewLogrus() (klog.Logger, func()) {
	logger := logrus.New()

	logger.SetLevel(logrus.DebugLevel)
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.JSONFormatter{
		DisableTimestamp:  true,
		DisableHTMLEscape: true,
		// PrettyPrint:       true,
	})
	return klogrus.NewLogger(logger), func() { log.Printf("\u001B[32m%s\u001B[0m\n", "logrus logger graceful close") }
}
