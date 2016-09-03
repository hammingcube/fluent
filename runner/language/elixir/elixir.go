package elixir

import (
	"github.com/maddyonline/fluent/runner/cmd"
	"github.com/maddyonline/fluent/runner/util"
	"path/filepath"
)

func Run(files []string, stdin string) (string, string, error) {
	workDir := filepath.Dir(files[0])
	sourceFiles := util.FilterByExtension(files, "ex")
	args := append([]string{"elixirc"}, sourceFiles...)
	return cmd.RunStdin(workDir, stdin, args...)
}
