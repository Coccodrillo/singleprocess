package singleprocess

import (
	"log"
	"testing"

	ps "github.com/mitchellh/go-ps"
	"github.com/stretchr/testify/assert"
)

type MockProcess struct {
	pid        int
	ppid       int
	executable string
}

func (m MockProcess) Pid() int {
	return m.pid
}

func (m MockProcess) PPid() int {
	return m.ppid
}

func (m MockProcess) Executable() string {
	return m.executable
}

func TestIsAnotherInstanceRunning(t *testing.T) {
	processName := "someprocessName"
	processes, err := ps.Processes()
	if err != nil {
		log.Fatalf("blah %v", err)
	}
	processes = append(processes, MockProcess{1, 2, processName})
	var countInstances int = 0
	for k := range processes {
		if processes[k].Executable() == processName {
			countInstances++
		}
	}

	assert.Equal(t, 1, countInstances, " The instance is running only once ")

	countInstances = 0
	processes = append(processes, MockProcess{2, 3, processName})
	for k := range processes {
		if processes[k].Executable() == processName {
			countInstances++
		}
	}
	assert.Equal(t, 2, countInstances, " The instance is running twice now ")
}
