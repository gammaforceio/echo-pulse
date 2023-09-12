package logger

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/coreos/go-systemd/journal"
)

func LogToFile(directory, filename, data string) {
	// Construct the full path for the file
	fullPath := filepath.Join(directory, filename)

	// Ensure the directory exists
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			journal.Print(journal.PriErr, fmt.Sprintf("Error creating directory: %s", err))
			return
		}
	}

	// Check if the file exists and create if not
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		_, err := os.Create(fullPath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			journal.Print(journal.PriErr, fmt.Sprintf("Error creating file: %s", err))
			return
		}
		fmt.Printf("File %s created.\n", fullPath)
	}

	// Write the data to the file
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		journal.Print(journal.PriErr, fmt.Sprintf("Error opening file: %s", err))
		return
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		journal.Print(journal.PriErr, fmt.Sprintf("Error writing to file: %s", err))
	}

	// Write the log to the systemd journal as well
	journal.Print(journal.PriInfo, data)
}
