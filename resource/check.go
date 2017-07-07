package resource

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/pivotal-topher-bullock/gist-resource/resource/internal"
)

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

func Check(request CheckRequest) ([]Version, error) {
	var versions []Version
	source := request.Source
	// TODO: use dis --> version := request.Version

	client := internal.NewGithubClient(source.Token)

	commits, _, err := client.Gists.ListCommits(context.Background(), source.Id, &github.ListOptions{})
	if err != nil {
		return versions, err
	}

	for _, commit := range commits {
		versions = append(versions, Version{"sha": *commit.Version})
	}

	return versions, nil
}
