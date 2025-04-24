package command

import (
	"github.com/urfave/cli/v2"
)

func (app *App) CmdResume(c *cli.Context) error {
	latestTimeEntry, err := app.client.GetLatestTimeEntry()
	if err != nil {
		return err
	}

	if latestTimeEntry != nil && latestTimeEntry.Stop == nil {
		return nil
	}

	if latestTimeEntry != nil {
		err := app.client.PutResumeTimeEntry(latestTimeEntry.WorkspaceID, latestTimeEntry.ID)
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
