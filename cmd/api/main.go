package main

import (
	"encoding/json"
	"flag"
	"github.com/pxcamus/go-sample-data-server/internal"
	"go.uber.org/zap"
	"net/http"
	"os"
	"strconv"
	"time"
)

const version = "1.0.0"

type application struct {
	config *internal.Config
	logger *zap.SugaredLogger
}

func main() {
	var cfg *internal.Config
	var env = flag.String("env", "dev", "Environment to use (dev|staging|production)")
	flag.Parse()
	cfg = internal.GetConfig(*env)

	logger := initializeLogger()
	if logger == nil {
		panic("logger is nil")
	}

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.ServerConfig.Port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.Infow("Starting server",
		"addr", srv.Addr,
		"env", cfg.ServerConfig.Env,
	)

	err := srv.ListenAndServe()
	app.logger.Error(err.Error())
	os.Exit(1)
}

func initializeLogger() *zap.SugaredLogger {
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "console",
	  "outputPaths": ["stdout"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelEncoder": "lowercase"
	  }
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.Sugar()
}
