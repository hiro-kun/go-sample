package cmd

// go run main.go resize hoge

import (
    "os"
    "io/ioutil"
    "strings"
    "github.com/spf13/cobra"
    "imageresize/imageresize/lib/resize"
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
          // TODO エラー処理後で修正
          panic(err)
        }

        var fileList[]string
        for _, file := range files {
          fileName := baseDir + file.Name()

          // TODO 一時的に.gitkeepを外す。根本解決は今度
          if strings.Contains(fileName, ".gitkeep") {
            continue
          }

          fileList = append(fileList, fileName)
        }

        d := resize.NewWorker()

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
