package log

import (
	"os"

	"go.uber.org/zap"
)

var Call *zap.SugaredLogger

func init() {
	logger, err := zap.NewProduction()
	if mode := os.Getenv("GIN_MODE"); mode != "release" {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		panic(err)
	}
	Call = logger.Sugar()
}
