package requirements

import (
	"fmt"

	"github.com/fujitsu-cf/cli/cf"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/models"
	"github.com/fujitsu-cf/cli/cf/terminal"
)

//go:generate counterfeiter -o fakes/fake_targeted_org_requirement.go . TargetedOrgRequirement
type TargetedOrgRequirement interface {
	Requirement
	GetOrganizationFields() models.OrganizationFields
}

type targetedOrgApiRequirement struct {
	ui     terminal.UI
	config core_config.Reader
}

func NewTargetedOrgRequirement(ui terminal.UI, config core_config.Reader) TargetedOrgRequirement {
	return targetedOrgApiRequirement{ui, config}
}

func (req targetedOrgApiRequirement) Execute() (success bool) {
	if !req.config.HasOrganization() {
		message := fmt.Sprintf(T("No org targeted, use '{{.Command}}' to target an org.", map[string]interface{}{"Command": terminal.CommandColor(cf.Name() + " target -o ORG")}))
		req.ui.Failed(message)
		return false
	}

	return true
}

func (req targetedOrgApiRequirement) GetOrganizationFields() (org models.OrganizationFields) {
	return req.config.OrganizationFields()
}
