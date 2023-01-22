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
        spent(args)
    case "made":
        made(args)
    case "create":
        create(args)
    case "account":
        account(args)
    default:
        fmt.Println("Unknown command")
    }
}
