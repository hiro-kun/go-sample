package cmd

// go run main.go resize hoge

import (
    "fmt"
    "os"
    "io/ioutil"
    "sync"
    "strings"

    "github.com/disintegration/imaging"
    "github.com/spf13/cobra"

    "imageresize/imageresize/lib/conf"
)

func init() {
    RootCmd.AddCommand(reizeCmd)
}

var reizeCmd = &cobra.Command{
    Use:   "resize",
    Long:  "resize cmd.",
    Run: func(cmd *cobra.Command, args []string) {

        baseDir := os.Getenv("GOPATH") + "/src" + conf.BaseImageDir

        files, err := ioutil.ReadDir(baseDir)
        if err != nil {
          panic(err)
        }

        var fileList[]string
        for _, file := range files {
          fileName := baseDir + file.Name()

          // 一時的に.gitkeepを外す。根本解決は今度
          if strings.Contains(fileName, ".gitkeep") {
            continue
          }

          fileList = append(fileList, fileName)
        }

        d := NewWorker()

        for _, file := range fileList {
          d.Add(file)
        }

        d.Start()

        d.Stop()

        /* プログラムの設計
             ※拡張子は指定しない。
             ※画像の一括変換に対応、ファイル単体指定は無し

           - go run main.go resize height width
        */
    },
}


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
					resize(str)
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

// 画像変換
func resize(filePath string) {
fmt.Println(filePath)

  f, _ := os.Open(filePath)
  defer f.Close()

  fi, err := f.Stat()
  if err != nil {
    return
  }

  fileName := fi.Name()
  fmt.Println(fileName)

  // load an image from file
  srcImage, err := imaging.Open(filePath)
  if err != nil {
    panic(err)
  }

  outDir := os.Getenv("GOPATH") + conf.OutImageDir

  // save the image to file
  // err = imaging.Save(srcImage, "/tmp/" + fileName)
  err = imaging.Save(srcImage, outDir + fileName)
  if err != nil {
    panic(err)
  }
}
