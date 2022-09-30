package utils

import (
	"fmt"
	"os"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
)

func printBanner() {
	isEnabled := true
	isColorEnabled := true
	reader, _ := os.Open("./banner.txt")
	banner.Init(colorable.NewColorableStdout(), isEnabled, isColorEnabled, reader)
	fmt.Println()
}
