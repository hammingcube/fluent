# Fluent -- a fork of prasmussen/glot-code-runner

Main differences: 
1. It is "go get"-able 
2. Adds a streaming mode

1. Default mode
```sh
cat example.json  | fluent
```
Output is produced after completion
```
{"stdout":"5\n2\n","stderr":"","error":""}

```

2. Stream mode
```sh

cat example.json  | fluent --stream
```
Output is streamed
```
5
2
```


Note: This is a hand-crafted fork and makes use of `go generate`.
