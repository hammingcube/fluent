package golang

//go:generate sed -n -e s/\.\.\/\.\./github.com\/maddyonline\/fluent\/runner/ -e "w golang.go" $GOPATH/src/github.com/prasmussen/glot-code-runner/language/golang/golang.go
