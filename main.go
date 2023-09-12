package main

import (
	//"flag"
	"fmt"

	"github.com/SiirRandall/echo-pulse/config"
	"github.com/SiirRandall/echo-pulse/logger"
	"github.com/SiirRandall/echo-pulse/server"
	flag "github.com/spf13/pflag"
)

func main() {
	// Define the flags
	var configPath string

	flag.StringVarP(&configPath, "config", "c", "/etc/echopulse/config.conf", "Path to config")

	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		logger.LogToFile("logs", "error.log", fmt.Sprintf("Error loading config: %s", err))
		fmt.Printf("Error loading config: %s", err) // Print error directly to console
		return
	}

	udpServer := server.NewUDPEchoServer(cfg.Log.Directory, cfg.Blacklist.Keywords)

	udpServer.Start(cfg.Server.IP, cfg.Server.Port)
}
