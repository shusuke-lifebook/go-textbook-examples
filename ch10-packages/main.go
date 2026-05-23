package main

import (
	"ch10-packages/internal/greeting"
	"fmt"
)

func main() {
	msg := greeting.Hello("Go")
	fmt.Println(msg)
}
