package service_test

import (
	"github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestService(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

type passingRequirement struct {
	Name string
}

func (r passingRequirement) Execute() bool {
	return true
}
