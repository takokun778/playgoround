package main

import "go.uber.org/zap"

func main() {
	plog, _ := zap.NewProduction()
	defer plog.Sync()

	plog.Debug("debug", zap.String("id", "1"))
	plog.Info("info", zap.String("id", "1"))
	plog.Warn("warn", zap.String("id", "1"))
	plog.Error("error", zap.String("id", "1"))

	dlog, _ := zap.NewDevelopment()
	defer dlog.Sync()

	dlog.Debug("debug", zap.String("id", "1"))
	dlog.Info("info", zap.String("id", "1"))
	dlog.Warn("warn", zap.String("id", "1"))
	dlog.Error("error", zap.String("id", "1"))
}
