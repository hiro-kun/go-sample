package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    // "github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
    Use:   "sample",
    Short: "This tool is pretty cool.",
    Long:  "This tool is a great convenience.",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func init() {
    cobra.OnInitialize()
    RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of sample",
    Long:  `All software has versions. This is sample`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("sample v1.0")
    },
}

func Sample(msg string)() {
  fmt.Println(msg)
}
