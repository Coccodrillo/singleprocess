package singleprocess

import (
	"log"
	"os"
	"runtime"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func IsAnotherInstanceRunning() bool {
	processName := strings.TrimLeft(os.Args[0], "./")
	var maxLength int
	switch runtime.GOOS {
	case "darwin":
		maxLength = 16
	case "linux":
		maxLength = 15
	}
	if maxLength > 0 && len(processName) > maxLength {
		processName = processName[:maxLength]
	}
	processes, err := ps.Processes()
	if err != nil {
		log.Fatalf("Error while retrieving process list %v", err)
	}
	var countInstances int
	for k := range processes {
		if processes[k].Executable() == processName {
			countInstances++
		}
	}
	return countInstances > 1
}
