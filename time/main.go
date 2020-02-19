package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UTC()
	since := time.Since(start)
	ms := int(since / time.Millisecond)
	fmt.Println(ms)
	fmt.Println(since)
}
