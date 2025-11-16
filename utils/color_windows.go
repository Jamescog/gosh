//go:build windows

package utils

import (
	"os"
	"unsafe"

	"golang.org/x/sys/windows"
)

type osVersionInfoEx struct {
	dwOSVersionInfoSize uint32
	dwMajorVersion      uint32
	dwMinorVersion      uint32
	dwBuildNumber       uint32
	dwPlatformId        uint32
	szCSDVersion        [128]uint16
}

func isWindows10OrAbove() bool {
	mod := windows.NewLazySystemDLL("ntdll.dll")
	proc := mod.NewProc("RtlGetVersion")

	var info osVersionInfoEx
	info.dwOSVersionInfoSize = uint32(unsafe.Sizeof(info))

	r, _, _ := proc.Call(uintptr(unsafe.Pointer(&info)))
	if r != 0 {
		return false
	}

	return info.dwMajorVersion >= 10
}

func enableVTProcessing() bool {
	if !isWindows10OrAbove() {
		return false
	}

	handle := windows.Handle(os.Stdout.Fd())
	var mode uint32
	if windows.GetConsoleMode(handle, &mode) != nil {
		return false
	}

	mode |= 0x0004
	return windows.SetConsoleMode(handle, mode) == nil
}
