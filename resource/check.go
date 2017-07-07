package resource

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/pivotal-topher-bullock/gist-resource/resource/internal"
)

type CheckRequest struct {
	Source  Source   `json:"source"`
	Version *Version `json:"version,omitempty"`
}

func Check(request CheckRequest) ([]Version, error) {
	source := request.Source

	client := internal.NewGithubClient(source.Token)

	commits, _, err := client.Gists.ListCommits(context.Background(), source.Id, &github.ListOptions{})
	if err != nil {
		return []Version{}, err
	}

	return versionsFrom(commits, request.Version), nil
}

func versionsFrom(commits []*github.GistCommit, from *Version) []Version {
	if from == nil {
		commit := commits[len(commits)-1]
		return []Version{Version{"sha": *commit.Version}}
	}

	var versions []Version

	fromVersion := *from

	for _, commit := range commits {
		version := Version{"sha": *commit.Version}
		versions = append([]Version{version}, versions...)

		if *commit.Version == fromVersion["sha"] {
			return versions
		}
	}

	return versions

}
