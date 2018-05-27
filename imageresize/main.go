/*

cd
git clone https://github.com/hiro-kun/go-sample.git
mkdir -p $GOPATH/src/github.com/hiro-kun/
ln -s /home/vagrant/go-sample/imageresize $GOPATH/src/github.com/hiro-kun/imageresize
cd $GOPATH/src/github.com/hiro-kun/imageresize
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
dep ensure

*/

package main

import (
  "fmt"
  "os"

  "github.com/hiro-kun/imageresize/cmd"
)

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
      fmt.Println(err)
      os.Exit(1)
  }
}
