package main

import (
	"fmt"
	"bosun.org/slog"
	"project/dup"
)

func main() {
    fmt.Println("Hello, go!!")
	
	slog.Info("Hello, slog!!")

	var words = []string{"hello", 
	"go", "hello"}
	fmt.Println(dup.GetDuplicates(words))
}
