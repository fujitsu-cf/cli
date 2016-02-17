package plugin_test

import (
	"path/filepath"

	"github.com/fujitsu-cf/cli/cf/commands/plugin"
	"github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/testhelpers/configuration"
	"github.com/fujitsu-cf/cli/testhelpers/plugin_builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPlugin(t *testing.T) {
	config := configuration.NewRepositoryWithDefaults()
	i18n.T = i18n.Init(config)

	_ = plugin.Plugins{}

	RegisterFailHandler(Fail)

	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_with_help")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_with_orgs")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_with_orgs_short_name")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_with_push")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_with_push_short_name")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_1")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "test_2")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "empty_plugin")
	plugin_builder.BuildTestBinary(filepath.Join("..", "..", "..", "fixtures", "plugins"), "alias_conflicts")

	RunSpecs(t, "Plugin Suite")
}
