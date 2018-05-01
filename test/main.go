package main

import (
    "errors"
)

func example(message string) (int, error) {
  if message == "hoge" {
       return 1, nil
  }
  return 0, errors.New("message must be hoge")
}
