package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	year := now.Year()

	fmt.Println(year)

	fmt.Println()

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)

	fmt.Println(fixed)
}