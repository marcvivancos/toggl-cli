package command

import (
	"github.com/marcvivancos/toggl-cli/cache"
	toggl "github.com/marcvivancos/toggl-cli/lib"
	"github.com/urfave/cli"
)

func (app *App) CmdStop(c *cli.Context) error {
	currentTimeEntry, err := app.client.GetCurrentTimeEntry()

	err = app.client.PutStopTimeEntry(currentTimeEntry.WorkspaceID, currentTimeEntry.ID)

	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(toggl.TimeEntry{})
	cache.Write()

	return nil
}
