// +build linux

package tcpinfo

import (
	"fmt"
	"net"
	"syscall"
	"unsafe"
)

type TCPInfo syscall.TCPInfo

func GetsockoptTCPInfo(tcpConn *net.TCPConn) (*TCPInfo, error) {
	if tcpConn == nil {
		return nil, fmt.Errorf("tcp conn is nil")
	}

	file, err := tcpConn.File()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fd := file.Fd()
	tcpInfo := TCPInfo{}
	size := unsafe.Sizeof(tcpInfo)
	_, _, errno := syscall.Syscall6(syscall.SYS_GETSOCKOPT, fd, syscall.SOL_TCP, syscall.TCP_INFO,
		uintptr(unsafe.Pointer(&tcpInfo)), uintptr(unsafe.Pointer(&size)), 0)
	if errno != 0 {
		return nil, fmt.Errorf("syscall failed. errno=%d", errno)
	}

	return &tcpInfo, nil
}
