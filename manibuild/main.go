package main

import (
  "fmt"
  "os"
  "path/filepath"
)

func main() {
  if len(os.Args) != 3 {
    _, _ = fmt.Fprintf(os.Stderr, "Invalid args. Expected an app build manifest path proceeded by a target github actions folder.")
    os.Exit(1)
  }

  manifestPath := filepath.Clean(os.Args[1])
  fmt.Println("Building", manifestPath)
  err := GenerateGithubActions(manifestPath, os.Args[2])
  if err != nil {
    _, _ = fmt.Fprintf(os.Stderr, err.Error())
    os.Exit(1)
  }
}
