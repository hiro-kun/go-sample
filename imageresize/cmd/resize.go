package cmd

// go run main.go resize hoge

import (
    "fmt"
    "os"
    "io/ioutil"

    "imageresize/imageresize/lib/conf"
    "github.com/spf13/cobra"
)

func init() {
    RootCmd.AddCommand(reizeCmd)
}

var reizeCmd = &cobra.Command{
    Use:   "resize",
    Long:  "resize cmd.",
    Run: func(cmd *cobra.Command, args []string) {
        baseDir := os.Getenv("GOPATH") + "/src" + conf.BaseImageDir
        outDir := os.Getenv("GOPATH") + "/src" + conf.OutImageDir

        files, err := ioutil.ReadDir(baseDir)
        if err != nil {
          panic(err)
        }

        for _, file := range files {
          fmt.Println(baseDir + file.Name())
        }

        fmt.Println(1111)
        fmt.Println(args)
        fmt.Println(baseDir)
        fmt.Println(outDir)



        /* プログラムの設計
             ※拡張子は指定しない。
             ※画像の一括変換に対応、ファイル単体指定は無し

           - go run main.go resize height width
        */
    },
}
