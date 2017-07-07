package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pivotal-topher-bullock/gist-resource/resource"
)

func main() {
	destination := os.Args[1]

	var input resource.InRequest

	err := json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := resource.In(destination, input)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.NewEncoder(os.Stdout).Encode(output)
	if err != nil {
		log.Fatalln(err)
	}
}
