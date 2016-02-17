package commands

import (
	"errors"

	"github.com/fujitsu-cf/cli/cf/api"
	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/flags"

	"github.com/fujitsu-cf/cli/cf/api/authentication"
	"github.com/fujitsu-cf/cli/cf/command_registry"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	"github.com/fujitsu-cf/cli/cf/requirements"
	"github.com/fujitsu-cf/cli/cf/terminal"
)

//go:generate counterfeiter -o fakes/fake_ssh_code_getter.go . SSHCodeGetter
type SSHCodeGetter interface {
	command_registry.Command
	Get() (string, error)
}

type OneTimeSSHCode struct {
	ui           terminal.UI
	config       core_config.ReadWriter
	authRepo     authentication.AuthenticationRepository
	endpointRepo api.EndpointRepository
}

func init() {
	command_registry.Register(&OneTimeSSHCode{})
}

func (cmd *OneTimeSSHCode) MetaData() command_registry.CommandMetadata {
	return command_registry.CommandMetadata{
		Name:        "ssh-code",
		Description: T("Get a one time password for ssh clients"),
		Usage:       T("CF_NAME ssh-code"),
	}
}

func (cmd *OneTimeSSHCode) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) ([]requirements.Requirement, error) {
	if len(fc.Args()) != 0 {
		cmd.ui.Failed(T("Incorrect Usage. No argument required\n\n") + command_registry.Commands.CommandUsage("ssh-code"))
	}

	reqs := append([]requirements.Requirement{}, requirementsFactory.NewApiEndpointRequirement())
	return reqs, nil
}

func (cmd *OneTimeSSHCode) SetDependency(deps command_registry.Dependency, _ bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	cmd.authRepo = deps.RepoLocator.GetAuthenticationRepository()
	cmd.endpointRepo = deps.RepoLocator.GetEndpointRepository()

	return cmd
}

func (cmd *OneTimeSSHCode) Execute(c flags.FlagContext) {
	code, err := cmd.Get()
	if err != nil {
		cmd.ui.Failed(err.Error())
	}

	cmd.ui.Say(code)
}

func (cmd *OneTimeSSHCode) Get() (string, error) {
	_, err := cmd.endpointRepo.UpdateEndpoint(cmd.config.ApiEndpoint())
	if err != nil {
		return "", errors.New(T("Error getting info from v2/info: ") + err.Error())
	}

	token, err := cmd.authRepo.RefreshAuthToken()
	if err != nil {
		return "", errors.New(T("Error refreshing oauth token: ") + err.Error())
	}

	return token, nil
}
