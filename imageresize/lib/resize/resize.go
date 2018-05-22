package resize

import(
  "sync"
  "os"
  "github.com/disintegration/imaging"
  "imageresize/imageresize/lib/conf"
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
	w := &Worker{
		queue: make(chan interface{}, queueLimit),
	}
	return w
}

// Add キューに処理を追加
func (w *Worker) Add(v interface{}) {
	w.queue <- v
}

// Start Worker開始
func (w *Worker) Start() {
	w.wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go func() {
			defer w.wg.Done()
			for v := range w.queue {
				if str, ok := v.(string); ok {
					resize(str)
				}
			}
		}()
	}
}

// Stop 停止処理
func (w *Worker) Stop() {
	close(w.queue)
	w.wg.Wait()
}

// 画像変換
func resize(filePath string) {

  f, _ := os.Open(filePath)
  defer f.Close()

  fi, err := f.Stat()
  if err != nil {
    // TODO エラー処理修正
    return
  }

  fileName := fi.Name()

  srcImage, err := imaging.Open(filePath)
  if err != nil {
    panic(err)
  }

  outDir := os.Getenv("GOPATH") + conf.OutImageDir

  err = imaging.Save(srcImage, outDir + fileName)
  if err != nil {
    panic(err)
  }
}
