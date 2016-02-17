package requirements

import (
	"fmt"

	"github.com/fujitsu-cf/cli/cf"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/terminal"
)

type TargetedSpaceRequirement struct {
	ui     terminal.UI
	config core_config.Reader
}

func NewTargetedSpaceRequirement(ui terminal.UI, config core_config.Reader) TargetedSpaceRequirement {
	return TargetedSpaceRequirement{ui, config}
}

func (req TargetedSpaceRequirement) Execute() (success bool) {
	if !req.config.HasOrganization() {
		message := fmt.Sprintf(T("No org and space targeted, use '{{.Command}}' to target an org and space", map[string]interface{}{"Command": terminal.CommandColor(cf.Name() + " target -o ORG -s SPACE")}))
		req.ui.Failed(message)
		return false
	}

	if !req.config.HasSpace() {
		message := fmt.Sprintf(T("No space targeted, use '{{.Command}}' to target a space", map[string]interface{}{"Command": terminal.CommandColor("cf target -s")}))
		req.ui.Failed(message)
		return false
	}

	return true
}
