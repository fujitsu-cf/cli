package requirements

import (
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	"github.com/fujitsu-cf/cli/cf/terminal"
)

type LoginRequirement struct {
	ui                     terminal.UI
	config                 core_config.Reader
	apiEndpointRequirement ApiEndpointRequirement
}

func NewLoginRequirement(ui terminal.UI, config core_config.Reader) LoginRequirement {
	return LoginRequirement{ui, config, ApiEndpointRequirement{ui, config}}
}

func (req LoginRequirement) Execute() (success bool) {
	if !req.apiEndpointRequirement.Execute() {
		return false
	}

	if !req.config.IsLoggedIn() {
		req.ui.Say(terminal.NotLoggedInText())
		return false
	}

	return true
}
