package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/marcvivancos/toggl-cli/cache"
	"github.com/marcvivancos/toggl-cli/command"
	toggl "github.com/marcvivancos/toggl-cli/lib"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

var (
	ConfigPath = os.Getenv("HOME")
)

const (
	ConfigName = ".toggl"
	ConfigType = "json"
)

const Name string = "toggl-cli"

var version string

func main() {
	cache.New(os.Getenv("HOME") + "/.toggl.cache.json")
	cache.Init()

	initialize()

	cmdApp := command.NewApp(viper.GetString("token"))

	app := cli.NewApp()
	app.Name = Name
	app.Version = version
	app.Usage = "Toggl API CLI Client"

	app.Flags = GlobalFlags
	app.Commands = []*cli.Command{
		{
			Name:   "start",
			Usage:  "Start time entry",
			Action: cmdApp.CmdStart,
			Flags: []cli.Flag{
				projectIDFlag,
				startTimeFlag,
				fromPreviousFlag,
			},
		},
		{
			Name:   "stop",
			Usage:  "End time entry",
			Action: cmdApp.CmdStop,
			Flags: []cli.Flag{
				endTimeFlag,
			},
		},
		{
			Name:   "resume",
			Usage:  "Resume time entry",
			Action: cmdApp.CmdResume,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "current",
			Usage:  "Show current time entry",
			Action: cmdApp.CmdCurrent,
			Flags:  []cli.Flag{},
		},
		{
			Name:   "workspaces",
			Usage:  "Show workspaces",
			Action: cmdApp.CmdWorkspaces,
		},
		{
			Name:   "projects",
			Usage:  "Show projects on current workspaces",
			Action: cmdApp.CmdProjects,
		},
		{
			Name:   "local",
			Usage:  "Set current dir workspace",
			Action: CmdLocal,
			Flags: []cli.Flag{
				projectIDFlag,
			},
		},
		{
			Name:   "global",
			Usage:  "Set global workspace",
			Action: CmdGlobal,
		},
	}
	app.CommandNotFound = CommandNotFound

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error: %s", err.Error())
		os.Exit(1)
	}
}

func requireToken() error {
	var token string
	var workspaces []toggl.Workspace
	var err error
	count := 0
	for count < 3 {
		fmt.Printf("Input API Token: ")
		fmt.Scan(&token)
		client := toggl.NewDefaultClient(token)
		workspaces, err = client.FetchWorkspaces()
		if err == nil {
			viper.Set("token", token)
			viper.Set("wid", workspaces[0].ID)
			return nil
		}
		count++
	}
	panic(fmt.Errorf("invalid token"))
}

func RootConfigFilePath() string {
	return filepath.Join(ConfigPath, ConfigName+"."+ConfigType)
}

func LocalConfigFilePath() string {
	return filepath.Join(".", ConfigName+"."+ConfigType)
}

func CreateConfig(path string, hash interface{}) {
	buf, _ := json.MarshalIndent(hash, "", "  ")
	err := os.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func LoadLocalConfig() error {
	localFilePath := LocalConfigFilePath()
	file, err := os.Open(localFilePath)
	if err != nil {
		return err
	}
	viper.MergeConfig(file)
	return nil
}

func initialize() {
	viper.SetConfigType(ConfigType)
	viper.SetConfigName(ConfigName)
	viper.AddConfigPath(ConfigPath)
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	LoadLocalConfig()

	if !viper.IsSet("token") {
		requireToken()
		CreateConfig(RootConfigFilePath(), viper.AllSettings())
	}
}
