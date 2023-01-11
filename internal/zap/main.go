package main

import "go.uber.org/zap"

func main() {
	plog, _ := zap.NewProduction()
	defer func() {
		if err := plog.Sync(); err != nil {
			panic(err)
		}
	}()

	plog.Debug("debug", zap.String("id", "1"))
	plog.Info("info", zap.String("id", "1"))
	plog.Warn("warn", zap.String("id", "1"))
	plog.Error("error", zap.String("id", "1"))

	dlog, _ := zap.NewDevelopment()
	defer func() {
		if err := dlog.Sync(); err != nil {
			panic(err)
		}
	}()

	dlog.Debug("debug", zap.String("id", "1"))
	dlog.Info("info", zap.String("id", "1"))
	dlog.Warn("warn", zap.String("id", "1"))
	dlog.Error("error", zap.String("id", "1"))
}
