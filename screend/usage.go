package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println(usageString())
}

func usageAndExit() {
	println(usageString())
	os.Exit(1)
}

func usageString() string {
	return (`Usage: screend [options] [command]

commands:
  help      show this help
  add       add and start a new process
  remove    remove a process
  list      list all processes
  start     start processes
  restart   restart processes
  stop      stop processes
  env       print environment variables
  set-env   set environment variables
  version   print version information

add:
  screend add [options] [command] [args...]

  options:
	-n, --name       name of the process
	-d, --dir        directory to start the process with
	-e, --env        environment variables to set (key=value, multiple allowed)
	-f, --force      allow update if process already exists
	-s, --start      start the process after

remove:
  screend remove [name] [options]

  options:
  --all        remove all processes

list:
  screend list [name] [options]

  options:
	--json        output as json

  If no name is given, all processes will be listed.
  
start:
  screend start [name]

  If no name is given, all processes will be started.

restart:
  screend restart [name]

  If no name is given, all processes will be restarted.

stop:
  screend stop [name]

  If no name is given, all processes will be stopped.

env:
  screend env [name]

  Print environment variables to given name process.
  
set-env:
  screend set-env [name] [key]=[value]

  If no name is given, the environment variable will be set for all processes.

  If value is not given, the environment variable will be unset.

version:
  screend version

  Print version information.
`)
}
