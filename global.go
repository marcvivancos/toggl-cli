package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
	cli "github.com/urfave/cli/v2"
)

func CmdGlobal(c *cli.Context) error {
	if !c.Args().Present() {
		return errors.New("command failed")
	}
	workspaceID, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return err
	}
	viper.Set("wid", workspaceID)

	CreateConfig(RootConfigFilePath(), viper.AllSettings())
	fmt.Printf("Global workspace set to: %d\n", workspaceID)

	return nil
}
