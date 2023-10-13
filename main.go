// This is the main package
package main

// Import necessary packages and libraries
import (
	"fmt" // For printing to console

	"github.com/gammaforceio/echo-pulse/config" // To load application's configuration settings
	"github.com/gammaforceio/echo-pulse/logger" // To log errors in a log file
	"github.com/gammaforceio/echo-pulse/server" // To manage server-related activities
	flag "github.com/spf13/pflag"               // Package pflag extends standard flag package with POSIX/GNU style flags
)

// Main function of the application
func main() {
	// Define the variable for the path of the config file
	var configPath string

	// Define command line flag for config file path. Default value is "/etc/echopulse/config.conf".
	flag.StringVarP(&configPath, "config", "c", "/etc/echopulse/config.conf", "Path to config")

	// Parse the defined flags from the command line.
	flag.Parse()

	// Load the configuration file
	cfg, err := config.LoadConfig(configPath)
	// If there is an error loading the config file,
	if err != nil {
		// log the error to a file
		logger.LogToFile("logs", "error.log", fmt.Sprintf("Error loading config: %s", err))
		// print the error directly to console
		fmt.Printf("Error loading config: %s", err)
		// Stop the execution as the config couldn't be loaded successfully.
		return
	}

	// Creates a new UDP echo server with provided log directory and blacklist keywords
	//	udpServer := server.NewUDPEchoServer(cfg.Log.Directory, cfg.Blacklist.Keywords)
	udpServer := server.NewEchoServer(cfg.Log.Directory, cfg.Blacklist.Keywords)
	// Start to listen on given IP and port
	udpServer.StartUDP(cfg.Server.IP, cfg.Server.Port)
	udpServer.StartTCP(cfg.Server.IP, cfg.Server.Port)
}
