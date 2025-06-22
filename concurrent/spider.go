package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task 任务表示需要爬取的URL
type Task struct {
	ID  int
	URL string
}

// Result 结果表示爬取结果
type Result struct {
	TaskID int
	Data   string
	Err    error
}

func main() {
	// 使用新的随机数生成器，替代已弃用的 rand.Seed
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	// 创建通道
	tasks := make(chan Task, 10)     // 任务队列
	results := make(chan Result, 10) // 结果队列
	done := make(chan struct{})      // 结束信号

	// 1. 启动任务分发器
	go func() {
		for i := 0; i < 5; i++ {
			tasks <- Task{
				ID:  i,
				URL: fmt.Sprintf("https://example.com/page/%d", i),
			}
			time.Sleep(time.Millisecond * 200) // 模拟任务生成间隔
		}
		close(tasks) // 任务分发完毕后关闭任务通道
	}()

	// 2. 启动3个工作协程处理任务
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, tasks, results, done, &wg, rng)
	}

	// 3. 结果收集器
	go func() {
		wg.Wait()      // 等待所有工作协程完成
		close(results) // 关闭结果通道
	}()

	// 4. 处理结果
	for res := range results {
		if res.Err != nil {
			fmt.Printf("任务#%d 失败: %v\n", res.TaskID, res.Err)
		} else {
			fmt.Printf("任务#%d 成功: %s\n", res.TaskID, res.Data)
		}
	}

	fmt.Println("所有任务处理完毕")
}

// 工作协程：处理任务并返回结果
func worker(id int, tasks <-chan Task, results chan<- Result, done <-chan struct{}, wg *sync.WaitGroup, rng *rand.Rand) {
	defer wg.Done()

	for {
		select {
		case task, ok := <-tasks:
			if !ok {
				fmt.Printf("工作协程#%d: 任务队列为空，退出\n", id)
				return
			}

			fmt.Printf("工作协程#%d: 开始处理任务#%d %s\n", id, task.ID, task.URL)

			// 使用新的随机数生成器
			// 模拟爬取过程（1-3秒随机耗时）
			time.Sleep(time.Duration(rng.Intn(3000)+1000) * time.Millisecond)

			// 模拟20%的失败率
			var res Result
			if rng.Intn(100) < 20 {
				res = Result{
					TaskID: task.ID,
					Err:    fmt.Errorf("模拟爬取失败"),
				}
			} else {
				res = Result{
					TaskID: task.ID,
					Data:   fmt.Sprintf("内容长度: %d", rng.Intn(1000)+500),
				}
			}

			// 将结果发送到结果通道，带超时处理
			select {
			case results <- res:
				fmt.Printf("工作协程#%d: 已提交任务#%d的结果\n", id, task.ID)
			case <-time.After(2 * time.Second):
				fmt.Printf("工作协程#%d: 提交任务#%d的结果超时\n", id, task.ID)
			case <-done:
				fmt.Printf("工作协程#%d: 收到结束信号，退出\n", id)
				return
			}

		case <-done:
			fmt.Printf("工作协程#%d: 收到结束信号，退出\n", id)
			return

		case <-time.After(5 * time.Second):
			fmt.Printf("工作协程#%d: 5秒无任务，进入空闲状态\n", id)
		}
	}
}
