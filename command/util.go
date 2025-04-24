package command

import (
	"encoding/csv"
	"os"

	"github.com/marcvivancos/toggl-cli/util"
	"github.com/urfave/cli/v2"
)

func NewWriter(c *cli.Context) (writer util.Writer) {
	if c.Bool("csv") {
		writer = csv.NewWriter(os.Stdout)
	} else {
		writer = util.NewTabWriter(os.Stdout)
	}
	return
}
