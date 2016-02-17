package requirements

import (
	"github.com/fujitsu-cf/cli/cf/api/spaces"
	"github.com/fujitsu-cf/cli/cf/models"
	"github.com/fujitsu-cf/cli/cf/terminal"
)

//go:generate counterfeiter -o fakes/fake_space_requirement.go . SpaceRequirement
type SpaceRequirement interface {
	Requirement
	SetSpaceName(string)
	GetSpace() models.Space
}

type spaceApiRequirement struct {
	name      string
	ui        terminal.UI
	spaceRepo spaces.SpaceRepository
	space     models.Space
}

func NewSpaceRequirement(name string, ui terminal.UI, sR spaces.SpaceRepository) *spaceApiRequirement {
	req := &spaceApiRequirement{}
	req.name = name
	req.ui = ui
	req.spaceRepo = sR
	return req
}

func (req *spaceApiRequirement) SetSpaceName(name string) {
	req.name = name
}

func (req *spaceApiRequirement) Execute() (success bool) {
	var apiErr error
	req.space, apiErr = req.spaceRepo.FindByName(req.name)

	if apiErr != nil {
		req.ui.Failed(apiErr.Error())
		return false
	}

	return true
}

func (req *spaceApiRequirement) GetSpace() models.Space {
	return req.space
}
