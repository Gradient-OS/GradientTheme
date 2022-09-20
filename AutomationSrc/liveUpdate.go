package main

import (
	"os/exec"
	"time"
)

func liveUpdate() {
	setAppSym("", true, false)
	exec.Command("cp", "./index.theme", tmpBuild+"/index.theme").Run()
	exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", "this-theme-does-not-exist").Run()
	exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", "SimplisticGradient").Run()
	time.Sleep(time.Second * 5)
}
