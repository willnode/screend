package main

func usage() {
	println(`Usage: screend [options] [command]

commands:
  help      show this help
  add       add and start a new process
  remove    remove a process
  list      list all processes
  ps        list all running processes
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
	-nr, --no-start  do not start the process after adding

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

  If no name is given, all environment variables will be printed.

set-env:
  screend set-env [name] [key]=[value]

  If no name is given, the environment variable will be set for all processes.

  If value is not given, the environment variable will be unset.

version:
  screend version

  Print version information.
`)
}
