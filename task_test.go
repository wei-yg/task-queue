package task_queue

import (
	"fmt"
	"testing"
	"time"
)

func addTasks(worker *Worker) {
	// 添加任务
	worker.Push(func() error {
		fmt.Println("Running task 1...")
		time.Sleep(1 * time.Second)                           // 模拟长时间运行的任务    1秒
		if err := worker.Sleep(5 * time.Second); err != nil { // 5秒   共6秒
			return err // 如果在 Sleep 期间接收到退出信号，返回错误
		}
		return nil
	})
	worker.Push(func() error {
		fmt.Println("Running task 2...")
		time.Sleep(1 * time.Second)                           // 模拟长时间运行的任务     1秒     共7秒
		if err := worker.Sleep(5 * time.Second); err != nil { //5秒   共13秒
			return err // 如果在 Sleep 期间接收到退出信号，也能返回错误
		}
		return nil
	})
	worker.Push(func() error {
		fmt.Println("Running task 3...")
		time.Sleep(10 * time.Second) // 模拟长时间运行的任务  // 10秒    共23秒
		return nil
	})
	fmt.Println("任务添加完成") // *** 注意这行输出的时机
}

// ******** 注意这里 建议在子协程添加任务，因为worker.Push采用的无缓冲通道，可能阻塞协程
func TestTask1(t *testing.T) {
	worker := NewWorker()
	go func() {
		addTasks(worker)
	}()

	time.Sleep(10 * time.Second) // 10秒停止
	worker.Stop()                // 停止执行
	time.Sleep(30 * time.Second) // 等待一段时间以便观察输出  等30秒 看是否再无输出
}

func TestTask2(t *testing.T) {
	// 此方法返回的是带缓冲的通道，需要你指定任务步数，任务步数需要 >= 你实际push的任务步数，不然也会有阻塞的问题
	worker := NewNoBlockingWorker(5)
	addTasks(worker)

	time.Sleep(10 * time.Second) // 10秒停止
	worker.Stop()                // 停止执行
	time.Sleep(30 * time.Second) // 等待一段时间以便观察输出  等30秒 看是否再无输出
}
