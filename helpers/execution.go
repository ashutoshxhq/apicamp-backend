package helpers

import (
	"bytes"
	"os/exec"
	"strings"
)

//ExecuteCommand ...
func ExecuteCommand(command string) {
	parts := strings.Split(command, " ")
	head := parts[0]
	args := parts[1:]
	cmd := exec.Command(head, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Run()
}

//ExecuteCommandInDirectory ...
func ExecuteCommandInDirectory(command string, folderPath string) {
	parts := strings.Split(command, " ")
	head := parts[0]
	args := parts[1:]
	cmd := exec.Command(head, args...)
	cmd.Dir = folderPath
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	cmd.Run()
}
