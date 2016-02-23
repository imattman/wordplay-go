package cli

import "github.com/codegangsta/cli"

func cmdWeb() cli.Command {
	return cli.Command{
		Name: "web",
		// Aliases: []string{"serve"},
		Usage:  "starts web application to serve requests",
		Action: actionWeb,
	}
}

func actionWeb(c *cli.Context) {

}
