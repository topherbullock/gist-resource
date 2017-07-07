package resource

import (
	"context"

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

type Files map[github.GistFilename]github.GistFile

func In(request InRequest) (Files, error) {
	var files Files
	source := request.Source
	version := request.Version
	client := internal.NewGithubClient(source.Token)

	gist, _, err := client.Gists.GetRevision(context.Background(), source.Id, version["sha"])
	if err != nil {
		return files, err
	}
	files = gist.Files

	if request.Params.Files != nil {
		files = Files{}

		for _, filename := range *request.Params.Files {
			files[github.GistFilename(filename)] = gist.Files[github.GistFilename(filename)]
		}
	}

	return gist.Files, nil
}
