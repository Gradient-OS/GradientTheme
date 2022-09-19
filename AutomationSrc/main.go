package main

import "os"

func main() {
	os.Remove("~/.icons/SimplisticGradient/icon-theme.cache")
	for {
		liveUpdate()
	}
}
