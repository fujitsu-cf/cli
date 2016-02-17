package commands_test

import (
	"github.com/fujitsu-cf/cli/cf/commands"
	"github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCommands(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	_ = commands.Api{}

	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}

type passingRequirement struct {
	Name string
}

func (r passingRequirement) Execute() bool {
	return true
}
