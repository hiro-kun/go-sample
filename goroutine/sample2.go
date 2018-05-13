package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	// Worker 構造体定義
	Worker struct {
		queue chan interface{}
		wg    sync.WaitGroup
	}
)

const (
	maxWorkers = 2
	queueLimit = 1000
)

// NewWorker NewWorker生成
func NewWorker() *Worker {
	d := &Worker{
		queue: make(chan interface{}, queueLimit),
	}
	return d
}

// Add キューに処理を追加
func (d *Worker) Add(v interface{}) {
	d.queue <- v
}

// Start Worker開始
func (d *Worker) Start() {
	d.wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go func() {
			defer d.wg.Done()
			for v := range d.queue {
				if str, ok := v.(string); ok {
					get(str)
				}
			}
		}()
	}
}

// Stop 停止処理
func (d *Worker) Stop() {
	close(d.queue)
	d.wg.Wait()
}

// get API call ダミー
func get(url string) {
	time.Sleep(2 * time.Second)
	fmt.Println(url)
}

func main() {
	apis := []string{
		"http://test.com/1",
		"http://test.com/2",
		"http://test.com/3",
		"http://test.com/4",
		"http://test.com/5",
	}

	d := NewWorker()

	for _, url := range apis {
		d.Add(url)
	}

	d.Start()

	d.Stop()
}
