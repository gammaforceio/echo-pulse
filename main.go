package main

import (
	"fmt"

	"github.com/gammaforceio/echo-pulse/config"
	"github.com/gammaforceio/echo-pulse/healthcheck"
	"github.com/gammaforceio/echo-pulse/logger"
	"github.com/gammaforceio/echo-pulse/server"
	flag "github.com/spf13/pflag"
)

func main() {
	var configPath string
	flag.StringVarP(&configPath, "config", "c", "/etc/echopulse/config.conf", "Path to config")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		logger.LogToFile("logs", "error.log", fmt.Sprintf("Error loading config: %s", err))
		fmt.Printf("Error loading config: %s", err)
		return
	}

	udpServer := server.NewUDPEchoServer(cfg.Log.Directory, cfg.Blacklist.Keywords)

	// Start health check HTTP server in a goroutine
	go healthcheck.Start("5060")

	// Start the UDP server
	udpServer.Start(cfg.Server.IP, cfg.Server.Port)
}
