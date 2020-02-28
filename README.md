# multimodule

Example of fundamental flaw of "local" packages

There is a standard way of running unit tests: ```go test ./...```

This command will run recursively ALL unit tests from all sub folders. But if you run it from root of repo
containing "local" package, you will get shit like this:

```service2/main.go:5:2: cannot find package "service2/localpackage" in any of:
        /usr/local/go/src/service2/localpackage (from $GOROOT)
        /home/akanaev/go/src/service2/localpackage (from $GOPATH)```
