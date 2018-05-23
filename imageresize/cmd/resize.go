package cmd

// go run main.go resize hoge

import (
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"

    "github.com/spf13/cobra"
    "github.com/go-ozzo/ozzo-validation"
    
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

        width, _ := strconv.Atoi(args[0])
        height, _ := strconv.Atoi(args[1])

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

        w.Start(width, height)

        w.Stop()

        /* プログラムの設計
             ※拡張子は指定しない。
             ※画像の一括変換に対応、ファイル単体指定は無し

           - go run main.go resize height width
        */
    },
}
