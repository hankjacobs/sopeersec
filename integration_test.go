package sopeersec

import (
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestUnconfined(t *testing.T) {
	td, err := ioutil.TempDir("", "sopeersec")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(td)

	socketPath := filepath.Join(td, "test.socket")
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()

	go func() {
		net.Dial("unix", socketPath)
	}()

	conn, _ := l.Accept()

	unixConn := conn.(*net.UnixConn)
	osFile, _ := unixConn.File()

	label, errno := GetsockoptPeerSec(int(osFile.Fd()), syscall.SOL_SOCKET)
	if errno != 0 {
		log.Fatal(errno)
	}

	expected := "unconfined"
	if label != expected {
		log.Fatalf("context was %v when it should have been %v", label, "expected")
	}
}
