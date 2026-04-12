# getting started

[link](https://go.dev/doc/tutorial/getting-started)

```bash
mkdir hello
cd hello
go mod init example/hello
touch hello.go
```

paste into [hello.go](./hello/hello.go):

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

run the program with

```bash
go run .
```