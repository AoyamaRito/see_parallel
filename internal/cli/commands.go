package cli

import (
	"fmt"
	"strings"

	"see_parallel/internal/analyzer"
	"see_parallel/internal/config"
	"see_parallel/internal/queue"
)

func AddToQueue(input []string) {
	if len(input) < 2 {
		fmt.Println("エラー: 質問とファイルを指定してください")
		fmt.Println("例: see_parallel queue '[\"質問\", \"file1.go\", \"file2.go\"]'")
		return
	}

	question := input[0]
	files := []string{}
	deep := false

	for i := 1; i < len(input); i++ {
		if i == len(input)-1 && input[i] == "deep" {
			deep = true
		} else {
			files = append(files, input[i])
		}
	}

	if len(files) == 0 {
		fmt.Println("エラー: 少なくとも1つのファイルを指定してください")
		return
	}

	task := queue.Task{
		Question: question,
		Files:    files,
		Deep:     deep,
	}

	q := queue.GetInstance()
	q.Add(task)

	fmt.Printf("キューに追加しました: [%d] %s\n", q.Size(), question)
	if deep {
		fmt.Println("  → 上位モデル(Gemini 2.0 Flash)を使用します")
	}
}

func ListQueue() {
	q := queue.GetInstance()
	tasks := q.GetAll()

	if len(tasks) == 0 {
		fmt.Println("キューは空です")
		return
	}

	fmt.Printf("キュー内のタスク数: %d\n\n", len(tasks))
	for i, task := range tasks {
		fmt.Printf("[%d] %s\n", i+1, task.Question)
		fmt.Printf("    ファイル: %s\n", strings.Join(task.Files, ", "))
		if task.Deep {
			fmt.Println("    モデル: Gemini 2.0 Flash (deep)")
		} else {
			fmt.Println("    モデル: Gemini 2.0 Flash Lite")
		}
		fmt.Println()
	}
}

func ClearQueue() {
	q := queue.GetInstance()
	size := q.Size()
	q.Clear()
	fmt.Printf("%d件のタスクをクリアしました\n", size)
}

func RunQueue(parallel int) {
	q := queue.GetInstance()
	tasks := q.GetAll()

	if len(tasks) == 0 {
		fmt.Println("キューは空です")
		return
	}

	context := config.GetContext()
	if context != "" {
		fmt.Printf("文脈情報: %s\n", context)
	}
	fmt.Printf("分析を開始します... (タスク数: %d, 並列数: %d)\n\n", len(tasks), parallel)

	results, duration := analyzer.RunParallel(tasks, parallel)

	fmt.Println("分析結果:\n")
	for i, result := range results {
		fmt.Printf("[%d] %s\n", i+1, result.Question)
		if result.Error != nil {
			fmt.Printf("→ エラー: %v\n", result.Error)
		} else {
			lines := strings.Split(result.Answer, "\n")
			for _, line := range lines {
				if strings.TrimSpace(line) != "" {
					fmt.Printf("→ %s\n", line)
				}
			}
		}
		fmt.Println()
	}

	fmt.Printf("実行時間: %.1f秒\n", duration.Seconds())
	fmt.Printf("並列数: %d\n", parallel)

	q.Clear()
}