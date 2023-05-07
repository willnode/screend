//go:generate go run git.rootprojects.org/root/go-gitver/v2
package main

import (
	"fmt"
	"os"
)

var (
	commit  = "0000000"
	version = "0.0.0-pre0+0000000"
	date    = "0000-00-00T00:00:00+0000"
)

func usage() {
	fmt.Printf(`Usage: screend [options] [command]

commands:
  help      show this help
  add       add and start a new process
  remove    remove a process
  list      list all processes
  ps        list all running processes
  start      start processes
  restart   restart processes
  stop      stop processes
  env       print environment variables
  set-env   set environment variables
  version   print version information

add:
  screend add [options] [command] [args...]

  options:
	-n, --name       name of the process
	-d, --dir        directory to start the command in
	-e, --env        environment variables to set (key=value, multiple allowed)
	-nr, --no-start  do not start the process after adding

remove:
  screend remove [name]

list:
  screend list [name] [options]

  options:
	--json        output as json

  If no name is given, all processes will be listed.
start:
  screend start [name]

  If no name is given, all processes will be start.

restart:
  screend restart [name]

  If no name is given, all processes will be restarted.

stop:
  screend stop [name]

  If no name is given, all processes will be stopped.

env:
  screend env [name]

  If no name is given, all environment variables will be printed.

set-env:
  screend set-env [name] [key]=[value]

  If no name is given, the environment variable will be set for all processes.

  If value is not given, the environment variable will be unset.

version:
  screend version
`)
}

func printVersion() {
	fmt.Printf(`screend %s
commit: %s
date: %s
`, version, commit, date)
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help":
		usage()
	case "add":
		add()
	case "remove":
		remove()
	case "list":
		list()
	case "start":
		start()
	case "restart":
		restart()
	case "stop":
		stop()
	case "env":
		env()
	case "set-env":
		setEnv()
	case "version":
		printVersion()
	default:
		usage()
		os.Exit(1)
	}
}
