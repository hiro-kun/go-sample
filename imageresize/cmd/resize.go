package cmd

// go run main.go resize hoge

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"

    "github.com/spf13/cobra"
    "github.com/go-ozzo/ozzo-validation"

    "imageresize/imageresize/lib/resize"
    "imageresize/imageresize/lib/conf"
)

func init() {
    RootCmd.AddCommand(reizeCmd)
}

type opt struct {
  width int
  height int
}

func (o opt) Validate() error {
	return validation.ValidateStruct(&o,
    // TODO エラー・メッセージを定数化
    validation.Field(&o.height, validation.Required, validation.Min(10).Error("error. height input between 10 to 500"), validation.Max(500).Error("error. height input between 10 to 500")),
    validation.Field(&o.height, validation.Required, validation.Min(10).Error("error. height input between 10 to 500"), validation.Max(500).Error("error. height input between 10 to 500")),
	)
}

var reizeCmd = &cobra.Command{
    Use:   "resize",
    Long:  "resize cmd.",
    Run: func(cmd *cobra.Command, args []string) {
      var opt = &opt{}

      opt.width, _ = strconv.Atoi(args[0])
      opt.height, _ = strconv.Atoi(args[1])

      if err := opt.Validate(); err != nil {
		      fmt.Println(err)
		      os.Exit(1)
	    }

      baseDir := os.Getenv("GOPATH") + conf.BaseImageDir

      files, err := ioutil.ReadDir(baseDir)
      if err != nil {
        fmt.Println("directory open error. %s", err)
        os.Exit(1)
      }

      var fileList[]string
      for _, file := range files {
        fileName := baseDir + file.Name()

        // TODO 一時的に.gitkeepを外す。根本解決は今度
        // 拡張子を特定の拡張子以外は弾くようにする
        if strings.Contains(fileName, ".gitkeep") {
          continue
        }

        fileList = append(fileList, fileName)
      }

      w := resize.NewWorker()
      for _, file := range fileList {
        w.Add(file)
      }

      w.Start(opt.width, opt.height)

      w.Stop()

        /* プログラムの設計
             ※拡張子は指定しない。
             ※画像の一括変換に対応、ファイル単体指定は無し

           - go run main.go resize height width
        */
    },
}
