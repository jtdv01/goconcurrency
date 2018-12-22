# What this repository is about

From O'Reilly course Hands on Go Concurrency:
https://learning.oreilly.com/videos/hands-on-concurrency-with/9781788993746

Reference on how to leverage Go's concurrency and patterns.

# Modules

## Intro
* Basics on `WaitGroups`
* 

# How to run examples
* The main file imports the chapters by module
* Relevant functions can be imported and run using `go run main.go`.
* Alternatively, the binary can be built using `make all`

# Note on Go Modules with GO111

* Go modules on 111 can be very opinionated on its project structure
* Create a directory in `~/go/src/github.com/<user>/<module>`
* In this case it is `~/go/src/github.com/jtdv01/goconcurrency`
* Init a module using:

```
GO111MODULE=on
go mod init
```

go.mod should now contain:

```
module github.com/jtdv01/goconcurrency
```