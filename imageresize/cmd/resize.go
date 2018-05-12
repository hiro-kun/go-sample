package cmd

// go run main.go resize hoge

import (
    "fmt"
    "os"
    "io/ioutil"

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

        // CPUコア数に応じてチャンネル数の上限を決める
        finished := make(chan bool)

        baseDir := os.Getenv("GOPATH") + "/src" + conf.BaseImageDir
        // outDir := os.Getenv("GOPATH") + "/src" + conf.OutImageDir

        files, err := ioutil.ReadDir(baseDir)
        if err != nil {
          panic(err)
        }

        var fileList[]string
        for _, file := range files {
          fileName := baseDir + file.Name()
          fileList = append(fileList, fileName)
          fmt.Println()
        }

        go func() {
          // load an image from file
          srcImage, err := imaging.Open(fileList[1])
          if err != nil {
            panic(err)
          }

          // save the image to file
          err = imaging.Save(srcImage, "/tmp/test.png")
          if err != nil {
            panic(err)
          }
        } ()

        // fmt.Println(args)
        // fmt.Println(baseDir)
        // fmt.Println(outDir)
        // fmt.Println(fileList)
        // fmt.Println(fileList[1])

        /*
           一つずつ取り出してresizeしていく。そこでgoルーチンを利用
           resize開始と終了でCUI上にデバッグ処理を出す
        */

        /* プログラムの設計
             ※拡張子は指定しない。
             ※画像の一括変換に対応、ファイル単体指定は無し

           - go run main.go resize height width
        */
    },
}
