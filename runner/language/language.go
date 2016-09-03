package language

import (
	"github.com/maddyonline/fluent/runner/language/assembly"
	"github.com/maddyonline/fluent/runner/language/ats"
	"github.com/maddyonline/fluent/runner/language/bash"
	"github.com/maddyonline/fluent/runner/language/c"
	"github.com/maddyonline/fluent/runner/language/clojure"
	"github.com/maddyonline/fluent/runner/language/coffeescript"
	"github.com/maddyonline/fluent/runner/language/cpp"
	"github.com/maddyonline/fluent/runner/language/csharp"
	"github.com/maddyonline/fluent/runner/language/d"
	"github.com/maddyonline/fluent/runner/language/elixir"
	"github.com/maddyonline/fluent/runner/language/elm"
	"github.com/maddyonline/fluent/runner/language/erlang"
	"github.com/maddyonline/fluent/runner/language/fsharp"
	"github.com/maddyonline/fluent/runner/language/golang"
	"github.com/maddyonline/fluent/runner/language/groovy"
	"github.com/maddyonline/fluent/runner/language/haskell"
	"github.com/maddyonline/fluent/runner/language/idris"
	"github.com/maddyonline/fluent/runner/language/java"
	"github.com/maddyonline/fluent/runner/language/javascript"
	"github.com/maddyonline/fluent/runner/language/julia"
	"github.com/maddyonline/fluent/runner/language/kotlin"
	"github.com/maddyonline/fluent/runner/language/lua"
	"github.com/maddyonline/fluent/runner/language/nim"
	"github.com/maddyonline/fluent/runner/language/ocaml"
	"github.com/maddyonline/fluent/runner/language/perl"
	"github.com/maddyonline/fluent/runner/language/perl6"
	"github.com/maddyonline/fluent/runner/language/php"
	"github.com/maddyonline/fluent/runner/language/python"
	"github.com/maddyonline/fluent/runner/language/ruby"
	"github.com/maddyonline/fluent/runner/language/rust"
	"github.com/maddyonline/fluent/runner/language/scala"
	"github.com/maddyonline/fluent/runner/language/swift"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"assembly":     assembly.Run,
	"ats":          ats.Run,
	"bash":         bash.Run,
	"c":            c.Run,
	"clojure":      clojure.Run,
	"coffeescript": coffeescript.Run,
	"csharp":       csharp.Run,
	"d":            d.Run,
	"elixir":       elixir.Run,
	"elm":          elm.Run,
	"cpp":          cpp.Run,
	"erlang":       erlang.Run,
	"fsharp":       fsharp.Run,
	"haskell":      haskell.Run,
	"idris":        idris.Run,
	"go":           golang.Run,
	"groovy":       groovy.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"julia":        julia.Run,
	"kotlin":       kotlin.Run,
	"lua":          lua.Run,
	"nim":          nim.Run,
	"ocaml":        ocaml.Run,
	"perl":         perl.Run,
	"perl6":        perl6.Run,
	"php":          php.Run,
	"python":       python.Run,
	"ruby":         ruby.Run,
	"rust":         rust.Run,
	"scala":        scala.Run,
	"swift":        swift.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
