package plugin_repo_test

import (
	"github.com/fujitsu-cf/cli/cf/commands/plugin_repo"
	"github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/testhelpers/configuration"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPluginRepo(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	_ = plugin_repo.RepoPlugins{}

	RegisterFailHandler(Fail)
	RunSpecs(t, "PluginRepo Suite")
}
