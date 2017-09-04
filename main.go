package main

import (
        "fmt"
	"time"
	"sync"
)

func main() {
    channel := make(chan string)
    var wg sync.WaitGroup
    wg.Add(1)

   go func() {
      channel <- "Hello from Channel!"
      println("Finishing goroutine")
       wg.Done()
   }()

    time.Sleep(time.Second)

    message := <- channel
    fmt.Println(message)
    wg.Wait()
}


