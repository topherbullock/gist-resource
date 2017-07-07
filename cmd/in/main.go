package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/pivotal-topher-bullock/gist-resource/resource"
)

func main() {
	destination := os.Args[1]

	var input resource.InRequest

	err := json.NewDecoder(os.Stdin).Decode(&input)
	if err != nil {
		log.Fatalln(err)
	}

	files, err := resource.In(input)
	if err != nil {
		log.Fatalln(err)
	}

	for filename, file := range files {
		data := []byte(*file.Content)
		err := ioutil.WriteFile(path.Join(destination, string(filename)), data, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
