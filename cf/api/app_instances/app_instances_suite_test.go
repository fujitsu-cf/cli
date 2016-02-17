package app_instances_test

import (
	"github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAppInstances(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	RegisterFailHandler(Fail)
	RunSpecs(t, "AppInstances Suite")
}
