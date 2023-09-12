// The config package holds configuration-related operations and structures
package config

// Import necessary packages
import (
	"github.com/BurntSushi/toml" // For TOML file operations
)

// Config is a struct that hold the configuration settings.
type Config struct {
	Server struct {
		IP   string // IP address for the server
		Port int    // Listening port for the server
	}
	Log struct {
		Directory string // Directory path where logs will be stored
	}
	Blacklist struct {
		Keywords []string // List of blacklisted words
	}
}

// LoadConfig reads a TOML file located at 'path' and decodes it into a Config object.
// It returns the pointer to the created Config structure if successful; otherwise, it returns an error.
func LoadConfig(path string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err // Return error if any issue happens during decoding the file
	}
	return &conf, nil // Return pointer to the filled Config structure
}
