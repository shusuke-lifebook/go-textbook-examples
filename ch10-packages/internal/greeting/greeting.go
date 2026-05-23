// Package greeting
package greeting

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", formatName(name))
}

// formatNameはパッケージ内部専用のヘルパー
func formatName(name string) string {
	return "[" + name + "]"
}
