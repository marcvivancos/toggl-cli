package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"strconv"

	cli "github.com/urfave/cli/v2"
)

func CmdLocal(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("command failed")
	}
	workspaceID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}

	var projectID int

	if c.IsSet("project-id") {
		projectID = c.Int("project-id")
	}

	path := filepath.Join(".", ConfigName+"."+ConfigType)
	CreateConfig(path, map[string]int{"wid": workspaceID, "pid": projectID})
	fmt.Printf("Local workspace set to: %d\n", workspaceID)

	return nil
}
