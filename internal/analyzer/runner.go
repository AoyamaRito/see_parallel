package analyzer

import (
	"context"
	"fmt"
	"sync"
	"time"

	"see_parallel/internal/fileutil"
	"see_parallel/internal/queue"
)

type Result struct {
	Index    int
	Question string
	Answer   string
	Error    error
}

func RunParallel(tasks []queue.Task, parallel int) ([]Result, time.Duration) {
	start := time.Now()
	ctx := context.Background()
	
	results := make([]Result, len(tasks))
	resultChan := make(chan Result, len(tasks))
	taskChan := make(chan struct {
		index int
		task  queue.Task
	}, len(tasks))

	for i, task := range tasks {
		taskChan <- struct {
			index int
			task  queue.Task
		}{index: i, task: task}
	}
	close(taskChan)

	var wg sync.WaitGroup
	for i := 0; i < parallel; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			client, err := NewGeminiClient(ctx)
			if err != nil {
				for taskData := range taskChan {
					resultChan <- Result{
						Index:    taskData.index,
						Question: taskData.task.Question,
						Error:    fmt.Errorf("failed to create Gemini client: %v", err),
					}
				}
				return
			}
			defer client.Close()

			for taskData := range taskChan {
				result := processTask(client, taskData.task)
				result.Index = taskData.index
				resultChan <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results[result.Index] = result
	}

	return results, time.Since(start)
}

func processTask(client *GeminiClient, task queue.Task) Result {
	expandedFiles, err := fileutil.ExpandFiles(task.Files)
	if err != nil {
		return Result{
			Question: task.Question,
			Error:    fmt.Errorf("failed to expand files: %v", err),
		}
	}

	content, err := fileutil.ReadAndCombineFiles(expandedFiles)
	if err != nil {
		return Result{
			Question: task.Question,
			Error:    fmt.Errorf("failed to read files: %v", err),
		}
	}

	answer, err := client.Analyze(task.Question, content, task.Deep)
	if err != nil {
		return Result{
			Question: task.Question,
			Error:    fmt.Errorf("failed to analyze: %v", err),
		}
	}

	return Result{
		Question: task.Question,
		Answer:   answer,
	}
}