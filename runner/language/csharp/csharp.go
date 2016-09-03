package csharp

import (
	"github.com/maddyonline/fluent/runner/cmd"
	"github.com/maddyonline/fluent/runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	binName := "a.exe"

	sourceFiles := util.FilterByExtension(files, "cs")
	args := append([]string{"mcs", "-out:" + binName}, sourceFiles...)
	stdout, stderr, err := cmd.Run(workDir, args...)
	if err != nil {
		return stdout, stderr, err
	}

	binPath := filepath.Join(workDir, binName)
	return cmd.RunStdin(workDir, stdin, "mono", binPath)
}
