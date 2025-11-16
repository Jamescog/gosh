//go:build !windows

package utils

func enableVTProcessing() bool {
	return true
}
