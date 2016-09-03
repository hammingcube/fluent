langlist="swift scala rust ruby python php perl6 perl ocaml nim lua kotlin julia javascript java idris haskell groovy golang fsharp erlang elm elixir d csharp cpp coffeescript clojure c bash ats assembly"
for arg in $langlist; do
	sed -n -e "s/prasmussen\/glot-code-runner/maddyonline\/fluent\/runner/"  -e "w runner/lang2/$arg/$arg.go" runner/language/$arg/$arg.go
done 
