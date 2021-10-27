package main

import (
    "flag"
    "log"
)

/**
 * @author Rancho
 * @date 2021/10/27
 */

func main() {
    var name string
    flag.StringVar(&name, "name", "Go语言编程之旅", "帮助信息")
    flag.StringVar(&name, "n", "Go语言编程之旅", "帮助信息")
    flag.Parse()

    log.Printf("name: %s", name)
}
