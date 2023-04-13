# Create a new executable

Before executing the next command, replace the version number with the specific version.
```go
 go build -o ./dist/kli -ldflags="-X 'github.com/konstellation-io/fake-kli/cmd.version=0.0.2'" cmd/kli.go
```
