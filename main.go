package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
    cmd := args[1]

    switch cmd {
    case "spent":
        fmt.Printf("Added new purchase of %s to %s category\n", args[2], args[3])
    case "made":
        fmt.Printf("Added new income (%s) from %s\n", args[2], args[3])
    case "create":
        fmt.Printf("Added new category (%s)\n", args[2])
    case "account":
        fmt.Println("Listing your purchases")
    default:
        fmt.Println("Unknown command")
    }
}
