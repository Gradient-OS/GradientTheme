package main

import "fmt"

func flags() {
	switch cmd.args[0] {
	case "-h":
		fmt.Println("-l: live update\n-b: bundle")
	case "-l":
		liveUpdate()
	case "-b":
		fmt.Println("not yet supported")
	default:
		buildThemeDev()
	}
}
