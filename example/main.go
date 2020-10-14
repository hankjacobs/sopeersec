package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"syscall"

	"github.com/hankjacobs/sopeersec"
)

func main() {
	l, err := net.Listen("unix", "/home/ubuntu/test/test1.socket")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()
	defer os.Remove("/home/ubuntu/test/test.socket")

	// Now do one of two things:
	// 1) use netcat or socat or something to connect to the above socket or
	// 2) bind mount the above directory into a docker container and do #1
	for {
		conn, _ := l.Accept()

		unixConn := conn.(*net.UnixConn)
		osFile, _ := unixConn.File()

		label, errno := sopeersec.GetsockoptPeerSec(int(osFile.Fd()), syscall.SOL_SOCKET)

		fmt.Println(errno, string(label))
		fmt.Println("---")
	}
}
