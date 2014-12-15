package main

import (
    "fmt"
    "time"
    "math/rand"
)

func boring(msg string) {
  for i := 0; ; i++ {
    fmt.Println("inside boring: ", msg, i)
    time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
  }
}

func boring2(msg string) {
  for i := 0; ; i++ {
    fmt.Println("inside boring2: ", msg, i)
    time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
  }
}

func main () {
  boring("boring!")
  boring2("boring!")
  fmt.Println("I'm listening.")
  time.Sleep(2 * time.Second)
  fmt.Println("You're boring; I'm leaving.")
}