- Initialize the go module
```
$ go mod init learn-go-with-tests
```

- Install go packages
```
$ go install golang.org/x/tools/cmd/godoc@latest
$ go install github.com/rakyll/gotest@latest
```

- Run go test files
```
$ gotest ./...
$ go test ./...
```

- Enable git hooks scripts.
```
$ git config --local core.hooksPath .githooks
```

- Check workflows status
```
$ gh workflow view Lint
$ gh workflow view Test
```