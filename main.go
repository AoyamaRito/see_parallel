package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"see_parallel/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "queue":
		handleQueueCommand()
	case "api":
		handleAPICommand()
	case "context":
		handleContextCommand()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func handleQueueCommand() {
	if len(os.Args) < 3 {
		printQueueUsage()
		os.Exit(1)
	}

	switch os.Args[2] {
	case "run":
		parallel := 10
		if len(os.Args) >= 5 && os.Args[3] == "--parallel" {
			p, err := strconv.Atoi(os.Args[4])
			if err != nil {
				fmt.Printf("Invalid parallel value: %s\n", os.Args[4])
				os.Exit(1)
			}
			parallel = p
		}
		cli.RunQueue(parallel)

	case "list":
		cli.ListQueue()

	case "clear":
		cli.ClearQueue()

	default:
		var input []string
		if err := json.Unmarshal([]byte(os.Args[2]), &input); err != nil {
			fmt.Printf("Invalid JSON format: %v\n", err)
			printQueueUsage()
			os.Exit(1)
		}
		cli.AddToQueue(input)
	}
}

func handleAPICommand() {
	if len(os.Args) < 4 || os.Args[2] != "set" {
		printAPIUsage()
		os.Exit(1)
	}

	apiKey := os.Args[3]
	cli.SetAPIKey(apiKey)
}

func handleContextCommand() {
	if len(os.Args) < 3 {
		printContextUsage()
		os.Exit(1)
	}

	switch os.Args[2] {
	case "set":
		if len(os.Args) < 4 {
			fmt.Println("エラー: 文脈情報を指定してください")
			os.Exit(1)
		}
		context := os.Args[3]
		cli.SetContext(context)
	case "get":
		cli.GetContext()
	case "clear":
		cli.ClearContext()
	default:
		printContextUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  see_parallel api set \"<API_KEY>\"")
	fmt.Println("  see_parallel context set \"文脈情報\"")
	fmt.Println("  see_parallel context get")
	fmt.Println("  see_parallel context clear")
	fmt.Println("  see_parallel queue '[\"質問\", \"ファイル1\", \"ファイル2\", ...]'")
	fmt.Println("  see_parallel queue run [--parallel N]")
	fmt.Println("  see_parallel queue list")
	fmt.Println("  see_parallel queue clear")
}

func printContextUsage() {
	fmt.Println("Context commands:")
	fmt.Println("  see_parallel context set \"文脈情報\"")
	fmt.Println("  see_parallel context get")
	fmt.Println("  see_parallel context clear")
}

func printAPIUsage() {
	fmt.Println("API commands:")
	fmt.Println("  see_parallel api set \"<API_KEY>\"")
}

func printQueueUsage() {
	fmt.Println("Queue commands:")
	fmt.Println("  see_parallel queue '[\"質問\", \"ファイル1\", \"ファイル2\", ...]'")
	fmt.Println("  see_parallel queue run [--parallel N]")
	fmt.Println("  see_parallel queue list")
	fmt.Println("  see_parallel queue clear")
}