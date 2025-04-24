# toggl CLI client
[toggl](https://toggl.com/) CLI Client, written in Golang.

## Description
[toggl](https://toggl.com/) is a time tracking web application.
This program will let you use the toggl in CLI.

## Usage
![demo image](https://cloud.githubusercontent.com/assets/6121271/21588531/0108bd18-d12b-11e6-9fdc-e65aa1f15768.gif)

```
NAME:
   toggl-cli - Toggl API CLI Client

USAGE:
   toggl-cli [global options] command [command options]

COMMANDS:
   start       Start time entry
   stop        End time entry
   resume      Resume time entry
   current     Show current time entry
   workspaces  Show workspaces
   projects    Show projects on current workspaces
   local       Set current dir workspace
   global      Set global workspace
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cache     (default: false)
   --csv       (default: false)
   --help, -h  show help
```

### Start commands
```
# start without project
toggl-cli start description
# start with project
toggl-cli -p 1234 description
# start from previous entry
toggl-cli start -fp
```

### Stop commands
```
# stop at the current time
toggl-cli stop
# stop at given time
toggl-cli stop -e 2025-04-23T13:00:00
```

## Install

To install, use `go get`:

```bash
$ go get github.com/marcvivancos/toggl-cli
```

### Register API token
When you run `toggl-cli` first time, you will be asked your toggl API token.
Please input toggl API token and register it.

## Acknowledgements
This project is a fork of [sachaos/toggl](https://github.com/sachaos/toggl) for personal use, I will implement some features but won't check them througly or make them "production ready" so I'm not merging them to the original project.
Many thanks to sachaos for the original implementation.
