package utils

import (
	"flag"
	"fmt"
	"os"
)

var Debug = false

func Initialize() {
	printBanner()
	var helpFlag = flag.Bool("help", false, "Show this help")
	flag.BoolVar(&Debug, "debug", false, "Enable Debugging")
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if Debug {
		fmt.Println("Debugging Enabled!!!")
	}

	ReadConfigs()
	InitServer()
}
