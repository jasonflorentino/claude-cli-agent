package main

import "fmt"

func Blue(str string) string {
	return fmt.Sprintf("\u001b[94m%s\u001b[0m", str)
}

func Gray(str string) string {
	return fmt.Sprintf("\u001b[2m%s\u001b[0m", str)
}

func Green(str string) string {
	return fmt.Sprintf("\u001b[92m%s\u001b[0m", str)
}

func Yellow(str string) string {
	return fmt.Sprintf("\u001b[93m%s\u001b[0m", str)
}
