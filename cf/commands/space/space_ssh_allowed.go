package space

import (
	"fmt"

	"github.com/fujitsu-cf/cli/cf/api/spaces"
	"github.com/fujitsu-cf/cli/cf/command_registry"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/requirements"
	"github.com/fujitsu-cf/cli/cf/terminal"
	"github.com/fujitsu-cf/cli/flags"
)

type SpaceSSHAllowed struct {
	ui        terminal.UI
	config    core_config.Reader
	spaceReq  requirements.SpaceRequirement
	spaceRepo spaces.SpaceRepository
}

func init() {
	command_registry.Register(&SpaceSSHAllowed{})
}

func (cmd *SpaceSSHAllowed) MetaData() command_registry.CommandMetadata {
	return command_registry.CommandMetadata{
		Name:        "space-ssh-allowed",
		Description: T("reports whether SSH is allowed in a space"),
		Usage:       T("CF_NAME space-ssh-allowed SPACE_NAME"),
	}
}

func (cmd *SpaceSSHAllowed) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires SPACE_NAME as argument\n\n") + command_registry.Commands.CommandUsage("space-ssh-allowed"))
	}

	cmd.spaceReq = requirementsFactory.NewSpaceRequirement(fc.Args()[0])
	reqs := []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
		requirementsFactory.NewTargetedOrgRequirement(),
		cmd.spaceReq,
	}

	return reqs, nil
}

func (cmd *SpaceSSHAllowed) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	return cmd
}

func (cmd *SpaceSSHAllowed) Execute(fc flags.FlagContext) {
	space := cmd.spaceReq.GetSpace()

	if space.AllowSSH {
		cmd.ui.Say(fmt.Sprintf(T("ssh support is enabled in space ")+"'%s'", space.Name))
	} else {
		cmd.ui.Say(fmt.Sprintf(T("ssh support is disabled in space ")+"'%s'", space.Name))
	}
}
