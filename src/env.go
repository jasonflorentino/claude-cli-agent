package main

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

func LoadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, found := strings.Cut(line, "=")
		if !found {
			continue
		}
		os.Setenv(key, value)
	}

	return scanner.Err()
}

var IgnoredFiles = make([]string, 2)

func IsFileIgnored(str string) bool {
	return slices.Contains(IgnoredFiles, str)
}

func LoadIgnoredFiles() error {
	file, err := os.Open(".gitignore")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		IgnoredFiles = append(IgnoredFiles, line)
	}

	return scanner.Err()

}
