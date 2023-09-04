package main

import (
	"github.com/hugiot/gioc"
	"github.com/hugiot/gioc-examples/ioc/provider"
	"github.com/hugiot/gioc-examples/ioc/service"
	"go.uber.org/zap"
)

func main() {
	gioc.AddServerProvider(&provider.AppServiceProvider{})
	gioc.Boot()

	logger := gioc.Make(service.Logger).(*zap.Logger)
	defer func() {
		_ = logger.Sync()
	}()

	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
