# go-tcpinfo

This is a small wrapper around the syscall library, so you don't have to mess around with it whenever you want to get tcp info for a connection

## Usage

```go
tcpInfo, err := tcpinfo.GetsockoptTCPInfo(&conn)
if err != nil {
    panic(err)
}
```

## Example

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"sync"

	"github.com/brucespang/go-tcpinfo"
)

func handleConn(conn net.Conn) {
	io.Copy(ioutil.Discard, conn)
}

func server(wg *sync.WaitGroup) {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	wg.Done()

	// accept connection on port
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	go io.Copy(ioutil.Discard, conn)

	_, err = conn.Write([]byte("hihihihihihihi"))
	if err != nil {
		panic(err)
	}

	tcpInfo, err := tcpinfo.GetsockoptTCPInfo(conn.(*net.TCPConn))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tcpInfo)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go server(&wg)
	wg.Wait()
	client()
}


```
