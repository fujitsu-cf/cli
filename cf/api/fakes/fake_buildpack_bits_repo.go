package fakes

import (
	"github.com/fujitsu-cf/cli/cf/errors"
	"github.com/fujitsu-cf/cli/cf/models"
)

type FakeBuildpackBitsRepository struct {
	UploadBuildpackErr         bool
	UploadBuildpackApiResponse error
	UploadBuildpackPath        string
}

func (repo *FakeBuildpackBitsRepository) UploadBuildpack(buildpack models.Buildpack, dir string) error {
	if repo.UploadBuildpackErr {
		return errors.New("Invalid buildpack")
	}

	repo.UploadBuildpackPath = dir
	return repo.UploadBuildpackApiResponse
}
