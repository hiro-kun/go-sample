/*
mkdir go-sample
cd go-sample
git clone https://github.com/hiro-kun/go-sample.git
ln -s /home/vagrant/go-sample/imageresize/ /home/vagrant/.gvm/pkgsets/go1.9.4/global/src/imageresize
mkdir $GOPATH/src
mkdir $GOPATH/src/imageresize
dep ensure

dep ensure -add github.com/spf13/cobra/cobra
*/

package main

import (
  "fmt"
  "os"

  // ここには$GOPATH(/home/vagrant/.gvm/pkgsets/go1.9.4/global)+src配下のフルパスを記載
  // ※後ほどパスを修正
  "imageresize/imageresize/cmd"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
      fmt.Println(err)
      os.Exit(-1)
  }
  cmd.Sample("tests")
}
