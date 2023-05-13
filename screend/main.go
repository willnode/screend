package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func version() string {
	if buildInfo, ok := debug.ReadBuildInfo(); ok && buildInfo.Main.Version != "(devel)" {
		return buildInfo.Main.Version
	}
	return "dev"
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help", "--help":
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
	case "version", "--version":
		fmt.Println(version())
	default:
		usageAndExit()
	}
}
