package route

import (
	"github.com/fujitsu-cf/cli/cf/api"
	"github.com/fujitsu-cf/cli/cf/command_registry"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/models"
	"github.com/fujitsu-cf/cli/cf/requirements"
	"github.com/fujitsu-cf/cli/cf/terminal"
	"github.com/fujitsu-cf/cli/flags"
	"github.com/fujitsu-cf/cli/flags/flag"
)

type DeleteOrphanedRoutes struct {
	ui        terminal.UI
	routeRepo api.RouteRepository
	config    core_config.Reader
}

func init() {
	command_registry.Register(&DeleteOrphanedRoutes{})
}

func (cmd *DeleteOrphanedRoutes) MetaData() command_registry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["f"] = &cliFlags.BoolFlag{ShortName: "f", Usage: T("Force deletion without confirmation")}

	return command_registry.CommandMetadata{
		Name:        "delete-orphaned-routes",
		Description: T("Delete all orphaned routes (e.g.: those that are not mapped to an app)"),
		Usage:       T("CF_NAME delete-orphaned-routes [-f]"),
		Flags:       fs,
	}
}

func (cmd *DeleteOrphanedRoutes) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) (reqs []requirements.Requirement, err error) {
	if len(fc.Args()) != 0 {
		cmd.ui.Failed(T("Incorrect Usage. No argument required\n\n") + command_registry.Commands.CommandUsage("delete-orphaned-routes"))
	}

	reqs = append(reqs, requirementsFactory.NewLoginRequirement())
	return
}

func (cmd *DeleteOrphanedRoutes) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	cmd.routeRepo = deps.RepoLocator.GetRouteRepository()
	return cmd
}

func (cmd *DeleteOrphanedRoutes) Execute(c flags.FlagContext) {
	force := c.Bool("f")
	if !force {
		response := cmd.ui.Confirm(T("Really delete orphaned routes?{{.Prompt}}",
			map[string]interface{}{"Prompt": terminal.PromptColor(">")}))

		if !response {
			return
		}
	}

	cmd.ui.Say(T("Getting routes as {{.Username}} ...\n",
		map[string]interface{}{"Username": terminal.EntityNameColor(cmd.config.Username())}))

	apiErr := cmd.routeRepo.ListRoutes(func(route models.Route) bool {

		if len(route.Apps) == 0 {
			cmd.ui.Say(T("Deleting route {{.Route}}...",
				map[string]interface{}{"Route": terminal.EntityNameColor(route.Host + "." + route.Domain.Name)}))
			apiErr := cmd.routeRepo.Delete(route.Guid)
			if apiErr != nil {
				cmd.ui.Failed(apiErr.Error())
				return false
			}
		}
		return true
	})

	if apiErr != nil {
		cmd.ui.Failed(T("Failed fetching routes.\n{{.Err}}", map[string]interface{}{"Err": apiErr.Error()}))
		return
	}
	cmd.ui.Ok()
}
