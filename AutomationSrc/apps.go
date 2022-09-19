package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const tmpBuild string = "./../SimplisticGradient"

func setAppSym(outDir string, liveUpdate bool, bundle bool) {

	content, err := ioutil.ReadFile("./apps.json")
	if err != nil {
		log.Fatal("Error opening ./../apps.json\n", err)
	}

	var payload map[string][]string

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error Whilst Unmarshalling Data\n", err)
	}

	exec.Command("\\cp", "-R", "./apps/.", tmpBuild+"/apps/").Run()
	for item, value := range payload {
		graphic := fmt.Sprintf("%s.svg", item)
		fmt.Println(graphic)
		for i := range value {
			os.Symlink(fmt.Sprintf("./%s", graphic), fmt.Sprintf("%s/apps/%s.svg", tmpBuild, value[i]))
		}
	}
}
