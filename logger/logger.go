// Logger package facilitates logging operations
package logger

// Import necessary packages and libraries
import (
	"fmt"           // For printing to console
	"os"            // To handle file-system operations
	"path/filepath" // For constructing filesystem paths

	"github.com/coreos/go-systemd/journal" // For logging to systemd journal
)

// LogToFile logs data to a specified file and system's journal.
func LogToFile(directory, filename, data string) {
	// Construct the full path for the file
	fullPath := filepath.Join(directory, filename)

	// Ensure the directory exists. If it doesn't, create it.
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			// On error, print message to console and log to system's journal.
			fmt.Println("Error creating directory:", err)
			journal.Print(journal.PriErr, fmt.Sprintf("Error creating directory: %s", err))
			return
		}
	}

	// Check if the file exists. If it doesn't, create it.
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		_, err := os.Create(fullPath)
		if err != nil {
			// On error, print message to console and log to system's journal.
			fmt.Println("Error creating file:", err)
			journal.Print(journal.PriErr, fmt.Sprintf("Error creating file: %s", err))
			return
		}
		fmt.Printf("File %s created.\n", fullPath)
	}

	// Open the file in append mode and write the data to it.
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// On error, print message to console and log to system's journal.
		fmt.Println("Error opening file:", err)
		journal.Print(journal.PriErr, fmt.Sprintf("Error opening file: %s", err))
		return
	}
	defer file.Close() // Close the file when function exits

	// Write data into the file.
	_, err = file.WriteString(data)
	if err != nil {
		// On error, print message to console and log to system's journal.
		fmt.Println("Error writing to file:", err)
		journal.Print(journal.PriErr, fmt.Sprintf("Error writing to file: %s", err))
	}

	// Log the data to the systemd journal as well
	journal.Print(journal.PriInfo, data)
}
