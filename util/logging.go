package util

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func GetInitializeLogger() (*zap.SugaredLogger, error) {
	log, err := zap.NewDevelopment()
	Logger = log.Sugar()
	Logger.Info("Initialize Zap Logger Successful!")
	return Logger, err
}
