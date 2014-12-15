package main

import (
    "fmt"
    "time"
    "math/rand"
    )

func boring(msg string, c chan string) {
  for i := 0; ; i++ {
    c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
    time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
  }
}

func main () {
  c := make(chan string)        // create the channel
  go boring("boring!", c)
  for i := 0; i < 5; i++ {
    // receiving from a channel
    // the "arrow" indicates the direction of data flow
    fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
  }
  fmt.Println("You're boring; I'm leaving.")
}