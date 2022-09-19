package main

import (
	"os/exec"
	"time"
)

func liveUpdate() {
	setAppSym("", true, false)
	exec.Command("\\cp", "./index.theme", tmpBuild+"/index.theme").Run()
	exec.Command("touch", "./../SimplisticGradient")
	exec.Command("update-icon-caches", "/usr/share/icons/*", "~/.icons/*").Run()
	time.Sleep(time.Minute)
}
