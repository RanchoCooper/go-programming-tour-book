package main

import (
    "log"

    "github.com/RanchoCooper/go-programming-tour-book/tour/cmd"
)

/**
 * @author Rancho
 * @date 2021/11/4
 */

func main() {
    err := cmd.Execute()
    if err != nil {
        log.Fatalf("cmd.Execute err: %v", err)
    }
}
