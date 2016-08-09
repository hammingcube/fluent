package runner

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/maddyonline/fluent/runner/cmd"
	"github.com/maddyonline/fluent/runner/language"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Payload struct {
	Language string          `json:"language"`
	Files    []*InMemoryFile `json:"files"`
	Stdin    string          `json:"stdin"`
	Command  string          `json:"command"`
}

type InMemoryFile struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Result struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

func Run() string {
	payload := &Payload{}
	err := json.NewDecoder(os.Stdin).Decode(payload)

	if err != nil {
		exitF("Failed to parse input json (%s)\n", err.Error())
	}

	// Ensure that we have at least one file
	if len(payload.Files) == 0 {
		exitF("No files given\n")
	}

	// Check if we support given language
	if !language.IsSupported(payload.Language) {
		exitF("Language '%s' is not supported\n", payload.Language)
	}

	// Write files to disk
	filepaths, err := writeFiles(payload.Files)
	if err != nil {
		exitF("Failed to write file to disk (%s)", err.Error())
	}

	var stdout, stderr string

	// Execute the given command or run the code with
	// the language runner if no command is given
	if payload.Command == "" {
		stdout, stderr, err = language.Run(payload.Language, filepaths, payload.Stdin)
	} else {
		workDir := filepath.Dir(filepaths[0])
		stdout, stderr, err = cmd.RunBashStdin(workDir, payload.Command, payload.Stdin)
	}
	return printResult(stdout, stderr, err)
}

// Writes files to disk, returns list of absolute filepaths
func writeFiles(files []*InMemoryFile) ([]string, error) {
	// Create temp dir
	tmpPath, err := ioutil.TempDir("", "")
	if err != nil {
		return nil, err
	}

	paths := make([]string, len(files), len(files))
	for i, file := range files {
		path, err := writeFile(tmpPath, file)
		if err != nil {
			return nil, err
		}

		paths[i] = path

	}
	return paths, nil
}

// Writes a single file to disk
func writeFile(basePath string, file *InMemoryFile) (string, error) {
	// Get absolute path to file inside basePath
	absPath := filepath.Join(basePath, file.Name)

	// Create all parent dirs
	err := os.MkdirAll(filepath.Dir(absPath), 0775)
	if err != nil {
		return "", err
	}

	// Write file to disk
	err = ioutil.WriteFile(absPath, []byte(file.Content), 0664)
	if err != nil {
		return "", err
	}

	// Return absolute path to file
	return absPath, nil
}

func exitF(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(1)
}

const DEFAULT_OUT_STR = `{"stdout": "", "stderr": "", "error": "%s"}`

func printResult(stdout, stderr string, err error) string {
	result := &Result{
		Stdout: stdout,
		Stderr: stderr,
		Error:  errToStr(err),
	}
	var outStr bytes.Buffer
	err = json.NewEncoder(&outStr).Encode(result)
	if err != nil {
		return fmt.Sprintf(DEFAULT_OUT_STR, err.Error())
	} else {
		return outStr.String()
	}
}

func errToStr(err error) string {
	if err != nil {
		return err.Error()
	}

	return ""
}
