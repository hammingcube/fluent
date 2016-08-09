package language

import (
	"github.com/maddyonline/fluent/runner/language/cpp"
	"github.com/maddyonline/fluent/runner/language/golang"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"cpp": cpp.Run,
	"go":  golang.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
