package sopeersec

import (
	"syscall"
	"unsafe"
)

// GetsockoptPeerSec invokes getsockopt with an optname of SO_PEERSEC
// returning the security state associated with the given socket file descriptor
func GetsockoptPeerSec(fd, level int) (string, syscall.Errno) {

	retried := false
	len := uint32(1024)
	val := make([]byte, len)

retry:
	errno := getsockopt(fd, level, syscall.SO_PEERSEC, unsafe.Pointer(&val[0]), &len)
	if errno == 0 {
		return string(val[:len]), 0
	}

	if errno == syscall.ERANGE && !retried {
		// ERANGE is returned when the array used for the label is too small
		// if that is the case, len contains the correct size for the label
		// so we try again.
		retried = true
		goto retry
	}

	return "", errno
}

func getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *uint32) (errno syscall.Errno) {
	_, _, errno = syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
	return
}
