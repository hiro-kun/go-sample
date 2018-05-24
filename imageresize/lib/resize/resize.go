package resize

import(
  "sync"
  "os"
  "fmt"
  "log"

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
func (w *Worker) Start(width int, height int) {
	w.wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go func() {
			defer w.wg.Done()
			for v := range w.queue {
				if filePath, ok := v.(string); ok {
					err := resize(width, height, filePath)
          if err != nil {
            log.Printf("[ERROR]: %s ", err)
          } else {
            fmt.Println(filePath + " is resized.")
          }
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

// resize 画像変換
func resize(w int, h int, filePath string) error {

  f, _ := os.Open(filePath)
  defer f.Close()

  fi, err := f.Stat()
  if err != nil {
    return fmt.Errorf(" %s is file open failed. %s", err)
  }

  fileName := fi.Name()
  outDir := os.Getenv("GOPATH") + conf.OutImageDir

  srcImage, err := imaging.Open(filePath)
  if err != nil {
    return fmt.Errorf(" %s is image open failed. %s", filePath, err)
  }

  resizeImage := imaging.Resize(srcImage, w, h, imaging.Lanczos)
  if err != nil {
    return fmt.Errorf(" %s is image resize failed. %s", filePath, err)
  }

  err = imaging.Save(resizeImage, outDir + fileName)
  if err != nil {
    return fmt.Errorf(" %s is image save failed. %s", filePath, err)
  }

  return nil
}
