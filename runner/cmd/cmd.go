package cmd

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

func Run(workDir string, args ...string) (string, string, error) {
	return RunStdin(workDir, "", args...)
}

var RunStdin func(string, string, ...string) (string, string, error)

func BlockRunStdin(workDir, stdin string, args ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = workDir
	cmd.Stdin = strings.NewReader(stdin)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func StreamRunStdin(workDir, stdin string, args ...string) (string, string, error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = workDir
	cmd.Stdin = strings.NewReader(stdin)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return "", "", err
}

func RunBash(workDir, command string) (string, string, error) {
	return Run(workDir, "bash", "-c", command)
}

func RunBashStdin(workDir, command, stdin string) (string, string, error) {
	return RunStdin(workDir, stdin, "bash", "-c", command)
}
