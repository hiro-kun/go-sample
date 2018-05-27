package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    // コマンド名
    Use:   "Image resize script.",
    // コマンドの説明
    Long:  "This tool will change the image size.",
    // スクリプト記載箇所
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
    cobra.OnInitialize()
    RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Long:  `All software has versions. This is Imageresize`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Imageresize v1.0")
    },
}
