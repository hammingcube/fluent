# Fluent -- a fork of [prasmussen/glot-code-runner](https://github.com/prasmussen/glot-code-runner)

### Note
This is a hand-crafted fork and makes use of `go generate`.

### Main differences:

1. It is "go get"-able
2. Adds a streaming mode


#### Default mode

```sh
cat example.json  | fluent
```
Output is produced after completion
```
{"stdout":"5\n2\n","stderr":"","error":""}

```

#### Stream mode

```sh

cat example.json  | fluent --stream
```
Output is streamed
```
5
2
```


# Misc Notes

```
docker run --rm -it -v $PWD:/go/bin -e GOPATH=/go -w /go/src/app golang go get -v github.com/maddyonline/fluent
```


