package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

var Model = anthropic.ModelClaude3_7SonnetLatest

func main() {
	if err := LoadIgnoredFiles(); err != nil {
		fmt.Printf("Error loading ignored files: %v\n", err)
		os.Exit(1)
	}
	if err := LoadEnv(); err != nil {
		fmt.Printf("Error loading .env: %v\n", err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			// fmt.Printf("Arg %d: %s\n", i+1, arg)
			if arg == "haiku" {
				Model = anthropic.ModelClaude3_5HaikuLatest
			}
		}
	}

	client := anthropic.NewClient(
		option.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
	)

	scanner := bufio.NewScanner(os.Stdin)
	getUserMessage := func() (string, bool) {
		if !scanner.Scan() {
			return "", false
		}
		return scanner.Text(), true
	}

	tools := []ToolDefinition{ReadFileDefinition, ListFilesDefinition, EditFileDefinition}
	agent := NewAgent(&client, getUserMessage, tools)

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")
	fmt.Println(Gray("Model: " + string(Model)))

	err := agent.Run(context.TODO())
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}
