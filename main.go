package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Welcome to the Go Text Editor!")
    
    if len(os.Args) < 2 {
        fmt.Println("Please provide a filename")
        os.Exit(1)
    }
    
    filename := os.Args[1]
    fmt.Printf("Editing file: %s\n", filename)

    contents, err := os.ReadFile(filename) 
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    fmt.Printf("\nContents of %s:\n%s", filename, string(contents)) 
}