# toggl CLI client
[toggl](https://toggl.com/) CLI Client, written in Golang.

## Description
[toggl](https://toggl.com/) is a time tracking web application.
This program will let you use the toggl in CLI.

## Usage
![demo image](https://cloud.githubusercontent.com/assets/6121271/21588531/0108bd18-d12b-11e6-9fdc-e65aa1f15768.gif)

```
$ toggl-cli --help
NAME:
   toggl-cli - Toggl API CLI Client

USAGE:
   toggl-cli [global options] command [command options] [arguments...]

VERSION:
   0.4.0

COMMANDS:
     start       Start time entry
     stop        End time entry
     restart     Restart previous running time entry
     list        Show todays time entry
     current     Show current time entry
     workspaces  Show workspaces
     projects    Show projects on current workspaces
     local       Set current dir workspace
     global      Set global workspace
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --cache
   --help, -h     show help
   --version, -v  print the version
```

### Start a time entry
Start a time entry with a project:
```
toggl-cli start test time entry --project-id 1234
```

Start a time entry from a given start time:
```
toggl-cli start test time entry [s|--start-time] 2023-09-15T14:00:00
```

Start a time entry from previous timer endtime:
``

### Restart previous time entry
Restarts the last running time entry
```
toggl-cli restart 
```

### Stop a time entry
Stop the current time entry:
```
toggl-cli stop
```

Stop the current time entry at a given time:
```
toggl-cli stop [-e|--end-time] 2023-09-15T14:30:00
```

### List time entries
List today time entries:
```
toggl-cli list
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
