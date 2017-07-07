package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pivotal-topher-bullock/gist-resource/resource"
)

func main() {
	var input resource.CheckRequest

	err := json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		log.Fatalln(err)
	}

	// TODO : actually validate the request
	versions, err := resource.Check(input)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalln(err)
	}
}
