package commands

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	. "github.com/fujitsu-cf/cli/cf/i18n"
	"github.com/fujitsu-cf/cli/cf/util"
	"github.com/fujitsu-cf/cli/flags"
	"github.com/fujitsu-cf/cli/flags/flag"

	"github.com/fujitsu-cf/cli/cf/api"
	"github.com/fujitsu-cf/cli/cf/command_registry"
	"github.com/fujitsu-cf/cli/cf/configuration/core_config"
	"github.com/fujitsu-cf/cli/cf/requirements"
	"github.com/fujitsu-cf/cli/cf/terminal"
	"github.com/fujitsu-cf/cli/cf/trace"
)

type Curl struct {
	ui       terminal.UI
	config   core_config.Reader
	curlRepo api.CurlRepository
}

func init() {
	command_registry.Register(&Curl{})
}

func (cmd *Curl) MetaData() command_registry.CommandMetadata {
	fs := make(map[string]flags.FlagSet)
	fs["i"] = &cliFlags.BoolFlag{ShortName: "i", Usage: T("Include response headers in the output")}
	fs["v"] = &cliFlags.BoolFlag{ShortName: "v", Usage: T("Enable CF_TRACE output for all requests and responses")}
	fs["X"] = &cliFlags.StringFlag{ShortName: "X", Usage: T("HTTP method (GET,POST,PUT,DELETE,etc)")}
	fs["H"] = &cliFlags.StringSliceFlag{ShortName: "H", Usage: T("Custom headers to include in the request, flag can be specified multiple times")}
	fs["d"] = &cliFlags.StringFlag{ShortName: "d", Usage: T("HTTP data to include in the request body, or '@' followed by a file name to read the data from")}
	fs["output"] = &cliFlags.StringFlag{Name: "output", Usage: T("Write curl body to FILE instead of stdout")}

	return command_registry.CommandMetadata{
		Name:        "curl",
		Description: T("Executes a request to the targeted API endpoint"),
		Usage: T(`CF_NAME curl PATH [-iv] [-X METHOD] [-H HEADER] [-d DATA] [--output FILE]

   By default 'CF_NAME curl' will perform a GET to the specified PATH. If data
   is provided via -d, a POST will be performed instead, and the Content-Type
   will be set to application/json. You may override headers with -H and the
   request method with -X.

   For API documentation, please visit http://apidocs.fujitsu-cf.org.

EXAMPLES:
   cf curl "/v2/apps" -X GET \
                      -H "Content-Type: application/x-www-form-urlencoded" \
                      -d 'q=name:myapp'

   cf curl "/v2/apps" -d @/path/to/file`),
		Flags: fs,
	}
}

func (cmd *Curl) Requirements(requirementsFactory requirements.Factory, fc flags.FlagContext) (reqs []requirements.Requirement, err error) {
	if len(fc.Args()) != 1 {
		cmd.ui.Failed(T("Incorrect Usage. An argument is missing or not correctly enclosed.\n\n") + command_registry.Commands.CommandUsage("curl"))
	}

	reqs = []requirements.Requirement{
		requirementsFactory.NewLoginRequirement(),
	}
	return
}

func (cmd *Curl) SetDependency(deps command_registry.Dependency, pluginCall bool) command_registry.Command {
	cmd.ui = deps.Ui
	cmd.config = deps.Config
	cmd.curlRepo = deps.RepoLocator.GetCurlRepository()
	return cmd
}

func (cmd *Curl) Execute(c flags.FlagContext) {
	path := c.Args()[0]
	method := c.String("X")
	headers := c.StringSlice("H")
	var body string
	if c.IsSet("d") {
		jsonBytes, err := util.GetContentsFromFlagValue(c.String("d"))
		if err != nil {
			cmd.ui.Failed(err.Error())
		}
		body = string(jsonBytes)
	}
	verbose := c.Bool("v")

	reqHeader := strings.Join(headers, "\n")

	if verbose {
		trace.EnableTrace()
	}

	responseHeader, responseBody, apiErr := cmd.curlRepo.Request(method, path, reqHeader, body)
	if apiErr != nil {
		cmd.ui.Failed(T("Error creating request:\n{{.Err}}", map[string]interface{}{"Err": apiErr.Error()}))
	}

	if verbose {
		return
	}

	if c.Bool("i") {
		cmd.ui.Say(responseHeader)
	}

	if c.String("output") != "" {
		err := cmd.writeToFile(responseBody, c.String("output"))
		if err != nil {
			cmd.ui.Failed(T("Error creating request:\n{{.Err}}", map[string]interface{}{"Err": err}))
		}
	} else {
		if strings.Contains(responseHeader, "application/json") {
			buffer := bytes.Buffer{}
			err := json.Indent(&buffer, []byte(responseBody), "", "   ")
			if err == nil {
				responseBody = buffer.String()
			}
		}

		cmd.ui.Say(responseBody)
	}
	return
}

func (cmd Curl) writeToFile(responseBody, filePath string) (err error) {
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(filePath), 0755)
	}

	if err != nil {
		return
	}

	return ioutil.WriteFile(filePath, []byte(responseBody), 0644)
}
