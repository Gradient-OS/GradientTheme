package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func setAppSym(outDir string, liveUpdate bool, bundle bool) int8 {

	content, err := ioutil.ReadFile("./../apps.json")
	if err != nil {
		log.Fatal("Error opening ./../apps.json\n", err)
	}

	var payload map[string]map[int]string

	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error Whilst Unmarshalling Data\n", err)
	}

	exec.Command("cp", "-R", "./apps", "./tempBuild")

	for item, value := range payload {
		graphic := fmt.Sprintf("%s.svg", item)
		for i, v := range value {
			os.Symlink(fmt.Sprintf("./tempbuild/%s", graphic), fmt.Sprintf("./tempbuild/%s.svg", v[i]))
		}
	}

	if liveUpdate == true {
		os.Rename("./tempbuild", outDir)
	} else if bundle == true && liveUpdate {
		return 0 // no errors
	}

	return 1 // An error occurred if it made it this far
}
