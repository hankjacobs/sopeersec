// +build !linux

package sopeersec

import "syscall"

func GetsockoptPeerSec(fd, level int) (string, syscall.Errno) {
	return "", syscall.EOPNOTSUPP
}
