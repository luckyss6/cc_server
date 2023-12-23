package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/luckyss6/cc_server/pkg/config"
	"github.com/luckyss6/cc_server/pkg/router"
	"github.com/luckyss6/cc_server/pkg/storage"
)

func Run() {
	// init config
	config.InitConfig()
	slog.Info("init config success")

	// init db
	storage.InitDB()

	slog.Info("init db success")

	// init server
	app := router.InitServer()
	go func() {
		slog.Info("server listen on :3000")
		if err := app.Listen(":3000"); err != nil {
			slog.Error(fmt.Sprintf("server listen error: %s", err.Error()))
			os.Exit(-1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	if err := app.Shutdown(); err != nil {
		slog.Error(fmt.Sprintf("server shutdown error: %s", err.Error()))
		os.Exit(-1)
	}

	slog.Info("server shutdown success")
}
