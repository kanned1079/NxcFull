package utils

import (
	"os"
	"unsafe"

	"syscall"
)

func GetTerminalWidthUnix() (int, error) {
	// 获取标准输出的文件描述符
	fd := int(os.Stdout.Fd())
	var dimensions [4]uint16

	// 调用 syscall.Syscall 进行 ioctl 系统调用，TIOCGWINSZ 用于获取窗口大小
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)))
	if errno != 0 {
		return 0, errno
	}
	// dimensions[1] 表示终端的宽度
	return int(dimensions[1]), nil
}
