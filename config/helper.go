package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/spf13/viper"
)

func GetSlogLevel() (slog.Level, error) {

	logLevelString := viper.GetString("logLevel")

	switch logLevelString {
	case "INFO":
		return slog.LevelInfo, nil
	case "WARN":
		return slog.LevelWarn, nil
	case "ERROR":
		return slog.LevelError, nil
	case "DEBUG":
		return slog.LevelDebug, nil
	}

	return slog.LevelInfo, fmt.Errorf("invalid log level: %s", logLevelString)
}

func NewLogger() (*slog.Logger, error) {

	logLevel, err := GetSlogLevel()
	if err != nil {
		return nil, err
	}

	logFormat := viper.GetString("logFormat")
	switch logFormat {
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})), nil
	case "text":
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel})), nil
	case "tint":
		return slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: logLevel, TimeFormat: time.Kitchen})), nil
	}

	return nil, fmt.Errorf("invalid log format: %s", logFormat)
}
