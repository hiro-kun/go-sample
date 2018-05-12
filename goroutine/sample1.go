package main

import (
    "fmt"
    "sync"
    "time"
)

func doSomething(u string) {
    fmt.Println(u)
    time.Sleep(2 * time.Second)
}

func main() {
    urls := []string{
        "http://test.com/1",
        "http://test.com/2",
        "http://test.com/3",
        "http://test.com/4",
        "http://test.com/5",
    }

    limit := make(chan struct{}, 2)

    var wg sync.WaitGroup
    for _, url := range urls {
        wg.Add(1)
        u := url
        go func() {
            limit <- struct{}{}
            defer wg.Done()
            doSomething(u)
            <-limit
        }()
        time.Sleep(200 * time.Millisecond)
    }
    wg.Wait()
}
