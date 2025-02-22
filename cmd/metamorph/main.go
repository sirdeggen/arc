package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/bitcoin-sv/arc/cmd"
	"github.com/ordishs/go-utils"
	"github.com/ordishs/gocore"
	"github.com/spf13/viper"
)

// Name used by build script for the binaries. (Please keep on single line)
const progname = "arc"

// // Version & commit strings injected at build with -ldflags -X...
var version string
var commit string

func init() {
	gocore.SetInfo(progname, version, commit)
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("failed to read config file config.yaml: %v \n", err)
		return
	}

	logLevel := viper.GetString("logLevel")
	logger := gocore.Log(progname, gocore.NewLogLevelFromString(logLevel))

	logger.Infof("VERSION\n-------\n%s (%s)\n\n", version, commit)

	go func() {
		profilerAddr := viper.GetString("metamorph.profilerAddr")
		if profilerAddr != "" {
			logger.Infof("Starting profile on http://%s/debug/pprof", profilerAddr)
			logger.Fatalf("%v", http.ListenAndServe(profilerAddr, nil))
		}
	}()

	shutdown, err := cmd.StartMetamorph(logger)
	if err != nil {
		logger.Fatalf("Error starting metamorph: %v", err)
	}

	// setup signal catching
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	<-signalChan

	appCleanup(logger, shutdown)
	os.Exit(1)
}

func appCleanup(logger utils.Logger, shutdown func()) {
	logger.Infof("Shutting down...")
	shutdown()
}
