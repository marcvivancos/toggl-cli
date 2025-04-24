package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var timelayout = "2006-01-02T15:04:05"

var GlobalFlags = []cli.Flag{
	&cli.BoolFlag{
		Name: "cache",
	},
	&cli.BoolFlag{
		Name: "csv",
	},
}

var projectIDFlag = &cli.IntFlag{
	Name:    "project-id",
	Aliases: []string{"P"},
	Usage:   "Project id",
}

var endTimeFlag = &cli.TimestampFlag{
	Name:     "end-time",
	Aliases:  []string{"e"},
	Usage:    "Stop time entry at a given time",
	Layout:   timelayout,
	Timezone: time.Local,
}

var startTimeFlag = &cli.TimestampFlag{
	Name:     "start-time",
	Aliases:  []string{"s"},
	Usage:    "Start time entry at a given time",
	Layout:   timelayout,
	Timezone: time.Local,
}

var fromPreviousFlag = &cli.BoolFlag{
	Name:    "from-previous",
	Aliases: []string{"fp"},
	Usage:   "Continue from previous time entry",
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
