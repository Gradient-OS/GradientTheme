package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

const tmpBuild string = "/home/ringwraith/.icons/SimplisticGradient" // ringwraith is user

func buildThemeDev() {
	setSymAll(true, false)

	// copy the theme declaration to the theme output
	exec.Command("cp", "./index.theme", tmpBuild+"/index.theme").Run()

	// reload SimplisticGradient theme
	exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", "this-theme-does-not-exist").Run()
	exec.Command("gsettings", "set", "org.gnome.desktop.interface", "icon-theme", "SimplisticGradient").Run()
}

func liveUpdate() {
	os.Remove("/home/ringwraith/.icons/SimplisticGradient/*")
	for {
		buildThemeDev()
		time.Sleep(time.Minute)
	}
}

func setSymAll(liveUpdate bool, bundle bool) {
	setSym("apps", true, false)
	setSym("locations", true, false)
}

func setSym(iconType string, liveUpdate bool, bundle bool) {
	var file string = fmt.Sprintf("./%s.json", iconType)
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error opening ./%s", file), err)
	}

	var payload map[string][]string

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error Whilst Unmarshalling Data\n", err)
	}

	exec.Command("cp", "-R", fmt.Sprintf("./%s/.", iconType), tmpBuild+fmt.Sprintf("/%s/", iconType)).Run()
	for item, value := range payload {
		graphic := fmt.Sprintf("%s.svg", item)
		for i := range value {
			os.Symlink(fmt.Sprintf("./%s", graphic), fmt.Sprintf("%s/%s/%s.svg", tmpBuild, iconType, value[i]))
		}
	}
}
