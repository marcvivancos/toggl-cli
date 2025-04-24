package command

import (
	"errors"
	"time"

	"github.com/marcvivancos/toggl-cli/cache"
	toggl "github.com/marcvivancos/toggl-cli/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

func (app *App) CmdStart(c *cli.Context) error {
	timeEntry := toggl.TimeEntry{}

	if !c.Args().Present() {
		return errors.New("command failed: description is required")
	}

	if c.Args().Present() {
		timeEntry.Description = c.Args().First()
	}
	timeEntry.WorkspaceID = viper.GetInt("wid")
	if c.IsSet("project-id") {
		timeEntry.ProjectID = c.Int("project-id")
	} else if viper.GetInt("pid") != 0 {
		timeEntry.ProjectID = viper.GetInt("pid")
	}

	if c.IsSet("start-time") {
		timeEntry.Start = c.Timestamp("start-time").UTC().Format(time.RFC3339)
	} else if c.IsSet("from-previous") {
		lastestTimeEntry, err := app.client.GetLatestTimeEntry()
		if err != nil {
			return err
		}
		if lastestTimeEntry == nil {
			timeEntry.Start = time.Now().UTC().Format(time.RFC3339)
		} else {
			if lastestTimeEntry.Stop == nil {
				app.client.PutStopTimeEntry(timeEntry.WorkspaceID, lastestTimeEntry.ID, time.Now().UTC())
				timeEntry.Start = time.Now().UTC().Format(time.RFC3339)
			} else {
				timeEntry.Start = *lastestTimeEntry.Stop
			}
		}
	} else {
		timeEntry.Start = time.Now().UTC().Format(time.RFC3339)
	}

	newTimeEntry, err := app.client.PostStartTimeEntry(timeEntry)
	if err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(newTimeEntry)
	cache.Write()

	return nil
}
