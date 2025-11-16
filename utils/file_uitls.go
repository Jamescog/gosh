package utils

import "fmt"

func HumanReadableSize(size int64) string {
	const (
		KB = 1 << (10 * 1)
		MB = 1 << (10 * 2)
		GB = 1 << (10 * 3)
		TB = 1 << (10 * 4)
	)
	switch {
	case size >= TB:
		return fmt.Sprintf("%.1fTB", float64(size)/float64(TB))
	case size >= GB:
		return fmt.Sprintf("%.1fGB", float64(size)/float64(GB))
	case size >= MB:
		return fmt.Sprintf("%.1fMB", float64(size)/float64(MB))
	case size >= KB:
		return fmt.Sprintf("%.1fKB", float64(size)/float64(KB))
	default:
		return fmt.Sprintf("%dB", size)
	}
}
