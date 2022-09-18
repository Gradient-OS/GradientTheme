package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func setAppSym() {
	content, err := ioutil.ReadFile("./../apps.json")
	if err != nil {
		log.Fatal("Error opening ./../apps.json\n", err)
	}

	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error Whilst Unmarshalling Data\n", err)
	}

	fmt.Printf("%s", payload["code"])
}
