package cmd

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

type opt struct {
  width int
  height int
}

func (o opt) Validate() error {
	return validation.ValidateStruct(&o,
    validation.Field(&o.height, validation.Required, validation.Min(10).Error(conf.ValidateInputErrorMessage), validation.Max(500).Error(conf.ValidateInputErrorMessage)),
    validation.Field(&o.height, validation.Required, validation.Min(10).Error(conf.ValidateInputErrorMessage), validation.Max(500).Error(conf.ValidateInputErrorMessage)),
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

        // .jpg,.png以外の拡張子の場合はスキップ
        if !strings.Contains(fileName, ".jpg") && !strings.Contains(fileName, ".png") {
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
    },
}
