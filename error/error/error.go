package error

import (
	"fmt"
	"os"
)

func Sample() error {
	// 存在しないファイルを読み込む
	_, err := os.Open("./sample.txt")
	if err != nil {
		return fmt.Errorf("File open failed. Error message is : %s", err)
	}

	return nil
}
