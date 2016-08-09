package main

import (
	"flag"
	"fmt"
	"github.com/maddyonline/fluent/runner"
	"github.com/maddyonline/fluent/runner/cmd"
)

func main() {
	var (
		stream = flag.Bool("stream", false, "stream stdout and stderr")
	)
	flag.Parse()
	if *stream {
		cmd.RunStdin = cmd.StreamRunStdin
		runner.Run()
	} else {
		cmd.RunStdin = cmd.BlockRunStdin
		fmt.Print(runner.Run())
	}
}
