package resource

import (
	"context"
	"io/ioutil"
	"log"
	"path"

	"github.com/google/go-github/github"
	"github.com/pivotal-topher-bullock/gist-resource/resource/internal"
)

type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
	Params  Params  `json:"params"`
}

type Params struct {
	Files *[]string `json:"files,omitempty"`
}

type InResult struct {
	Version  Version     `json:"version"`
	Metadata github.Gist `json:"metadata"`
}

type Files map[github.GistFilename]github.GistFile

func In(destination string, request InRequest) (InResult, error) {
	var files Files
	source := request.Source
	version := request.Version
	client := internal.NewGithubClient(source.Token)

	gist, _, err := client.Gists.GetRevision(context.Background(), source.Id, version["sha"])
	if err != nil {
		return InResult{}, err
	}
	files = gist.Files

	if request.Params.Files != nil {
		files = Files{}

		for _, filename := range *request.Params.Files {
			files[github.GistFilename(filename)] = gist.Files[github.GistFilename(filename)]
		}
	}

	for filename, file := range files {
		data := []byte(*file.Content)
		err := ioutil.WriteFile(path.Join(destination, string(filename)), data, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	gist.Files = nil

	return InResult{
		Version:  version,
		Metadata: *gist,
	}, nil
}
