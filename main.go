package main

import (
	"fmt"
	"tmc/TMC"
)

func main() {
	fmt.Println("hello")
	TMC.Init_uart("/dev/serial0", 32, 32)
	return
}
