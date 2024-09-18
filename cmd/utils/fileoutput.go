package utils

import (
	"fmt"
	"os"
)

func WriteToFile(scanPayload []byte) error {
	logFile, err := os.OpenFile("scanresult.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Writing to file: %v", err)
	}
	defer logFile.Close()

	_, err = logFile.Write(append(scanPayload, '\n'))
	if err != nil {
		return fmt.Errorf("Failed to write to file: %v", err)
	}
	return nil
}
