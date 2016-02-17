package commands

import (
	"github.com/fujitsu-cf/cli/cf/api/stacks"
	"github.com/fujitsu-cf/cli/cf/command_registry"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/requirements"
	"github.com/fujitsu-cf/cli/cf/terminal"
	"github.com/fujitsu-cf/cli/flags"
	"github.com/fujitsu-cf/cli/flags/flag"
)

type ListStack struct {
	ui         terminal.UI
	config     core_config.Reader
	stacksRepo stacks.StackRepository
}

func init() {
	command_registry.Register(&ListStack{})
}

func (cmd *ListStack) MetaData() command_registry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["guid"] = &cliFlags.BoolFlag{Name: "guid", Usage: T("Retrieve and display the given stack's guid. All other output for the stack is suppressed.")}

	return command_registry.CommandMetadata{
		Name:        "stack",
		Description: T("Show information for a stack (a stack is a pre-built file system, including an operating system, that can run apps)"),
		Usage:       T("CF_NAME stack STACK_NAME"),
		Flags:       fs,
		TotalArgs:   1,
	}
}

func (cmd *ListStack) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) (reqs []requirements.Requirement, err error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. Requires stack name as argument\n\n") + command_registry.Commands.CommandUsage("stack"))
	}

	reqs = append(reqs, requirementsFactory.NewLoginRequirement())
	return
}

func (cmd *ListStack) SetDependency(deps command_registry.Dependency, _ bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	cmd.stacksRepo = deps.RepoLocator.GetStackRepository()
	return cmd
}

func (cmd *ListStack) Execute(c flags.FlagContext) {
	stackName := c.Args()[0]

	stack, apiErr := cmd.stacksRepo.FindByName(stackName)

	if c.Bool("guid") {
		cmd.ui.Say(stack.Guid)
	} else {
		if apiErr != nil {
			cmd.ui.Failed(apiErr.Error())
			return
		}

		cmd.ui.Say(T("Getting stack '{{.Stack}}' in org {{.OrganizationName}} / space {{.SpaceName}} as {{.Username}}...",
			map[string]interface{}{"Stack": stackName,
				"OrganizationName": terminal.EntityNameColor(cmd.config.OrganizationFields().Name),
				"SpaceName":        terminal.EntityNameColor(cmd.config.SpaceFields().Name),
				"Username":         terminal.EntityNameColor(cmd.config.Username())}))

		cmd.ui.Ok()
		cmd.ui.Say("")

		table := terminal.NewTable(cmd.ui, []string{T("name"), T("description")})
		table.Add(stack.Name, stack.Description)
		table.Print()
	}
}
