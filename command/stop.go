package command

import (
	"fmt"
	"time"

	"github.com/marcvivancos/toggl-cli/cache"
	toggl "github.com/marcvivancos/toggl-cli/lib"
	"github.com/urfave/cli/v2"
)

func (app *App) CmdStop(c *cli.Context) error {
	currentTimeEntry, err := app.client.GetCurrentTimeEntry()
	if err != nil {
		return err
	}

	stopTime := time.Now()

	if c.IsSet("end-time") {
		stopTime = c.Timestamp("end-time").UTC()

		startTime, err := currentTimeEntry.ParsedStart()
		if err != nil {
			return fmt.Errorf("failed to parse start time: %v", err)
		}

		if !startTime.IsZero() && stopTime.Before(startTime) {
			return fmt.Errorf("end-time cannot be before start-time")
		}
	}

	if err := app.client.PutStopTimeEntry(currentTimeEntry.WorkspaceID, currentTimeEntry.ID, stopTime); err != nil {
		return err
	}

	cache.SetCurrentTimeEntry(toggl.TimeEntry{})
	cache.Write()

	return nil
}
