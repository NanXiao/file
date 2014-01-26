# File [![GoDoc](https://godoc.org/github.com/fatih/file?status.png)](http://godoc.org/github.com/fatih/file)

File includes several useful file related helper functions. It's not finished yet.

For usage see examples below or click on the godoc badge.

## Install

```bash
go get github.com/fatih/file
```

## Examples

```go
// copy recursively exampleDir to a new test directory
err := file.Copy("exampleDir", "test")

// copy a file into a folder
err := file.Copy("hello.txt", "./exampleDir")

// create a copy of a given file
err := file.Copy("hello.txt", "another.txt")

```