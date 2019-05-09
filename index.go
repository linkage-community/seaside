package main

import (
  "fmt"
  "github.com/otofune/seaside/config"
  "os"
)

func main () {
  c, err := config.LoadConfig()
  if err != nil {
    fmt.Printf("Can't load LoadConfig. %v\n", err)
    os.Exit(1)
  }
  fmt.Println(c)
}
