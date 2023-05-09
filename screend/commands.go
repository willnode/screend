package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/willnode/screend"
)

func add() {
	var p screend.Process
	args := os.Args[2:]
	p.Env = make(map[string]string)
	start := false
	force := false
	for _, str := range os.Environ() {
		splits := strings.SplitN(str, "=", 2)
		p.Env[splits[0]] = splits[1]
	}
	i := 0
	for i < len(args) {
		k := args[i]
		v := ""
		if i+1 < len(args) {
			v = args[i+1]
		}
		if p.Command == "" {
			if strings.HasPrefix(k, "-") {
				switch k {
				case "-n", "--name":
					i++
					if v == "" {
						log.Fatal("Empty option value: ", k)
					}
					if strings.ContainsAny(v, "=\t\n\v\f\r") {
						log.Fatal("Invalid option: ", v)
					}
					p.Name = v
				case "-d", "--dir":
					i++
					if v == "" {
						log.Fatal("Empty option value: ", k)
					}
					p.Dir = v
				case "-e", "--env":
					i++
					if v == "" {
						log.Fatal("Empty option value: ", k)
					}
					splits := strings.SplitN(v, "=", 2)
					if len(splits) != 2 {
						log.Fatal("Invalid environment variable: ", v)
					}
					if splits[1] == "" {
						delete(p.Env, splits[0])
					} else {
						p.Env[splits[0]] = splits[1]
					}

				case "-f", "--force":
					force = true
				case "-s", "--start":
					start = true
				default:
					log.Fatal("Invalid option: ", v)
				}
			} else {
				if p.Name == "" {
					// get file name
					p.Name = filepath.Base(k)
					if p.Name == "" {
						log.Fatal("Invalid command: ", k)
					}
				}
				p.Command = k
			}
		} else {
			p.Args = append(p.Args, k)
		}
		i++
	}

	if p.Command == "" {
		log.Fatal("Missing command")
	}

	if p.Dir == "" {
		p.Dir = p.Env["PWD"]
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	err = s.AddProcess(p, force)
	if err != nil {
		log.Fatal(err)
	}
	if start {
		if force {
			err = s.RestartProcess(p.Name)
		} else {
			err = s.StartProcess(p.Name)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func remove() {
	name := ""
	if len(os.Args) == 3 {
		name = os.Args[2]
	} else {
		log.Fatal("Missing name, use --all to remove all processes")
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	err = s.StopProcess(name)
	if err != nil {
		log.Fatal(err)
	}
	err = s.RemoveProcess(name)
	if err != nil {
		log.Fatal(err)
	}
}

func list() {
	name := ""
	format := "table"
	args := os.Args[2:]
	i := 0
	for i < len(args) {
		k := args[i]
		v := ""
		if i+1 < len(args) {
			v = args[i+1]
		}
		i++
		if strings.HasPrefix(k, "-") {
			switch k {
			case "-f", "--format":
				format = v
				i++
			default:
				log.Fatal("Invalid option: ", v)
			}
		} else {
			if name != "" {
				log.Fatal("Invalid argument: ", k)
			}
			name = k
		}
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	lists, err := s.ListProcesses(true, name)
	if err != nil {
		log.Fatal(err)
	}

	switch format {
	case "json":
		for _, v := range lists {
			line := struct {
				ScreenID string   `json:"screen_id"`
				Name     string   `json:"name"`
				Command  string   `json:"command"`
				Args     []string `json:"args"`
			}{
				v.ScreenID,
				v.Name,
				v.Command,
				v.Args,
			}

			str, err := json.Marshal(line)
			if err != nil {
				log.Fatal(err)
			}

			println(string(str))
		}
	case "table":
		w := tabwriter.NewWriter(os.Stdout, 10, 1, 1, ' ', 0)
		w.Write([]byte("ScreenID\tName\tCommand\tArgs\t\n"))
		for _, v := range lists {
			args, err := json.Marshal(v.Args)
			if err != nil {
				log.Fatal(err)
			}
			w.Write([]byte(strings.Join([]string{v.ScreenID, v.Name, v.Command, string(args)}, "\t")))
			w.Write([]byte("\n"))
		}
		w.Flush()
	default:
		log.Fatal("Invalid format: ", format)
	}
}

func start() {
	name := ""
	if len(os.Args) == 3 {
		name = os.Args[2]
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	err = s.StartProcess(name)
	if err != nil {
		log.Fatal(err)
	}
}

func stop() {
	name := ""
	if len(os.Args) == 3 {
		name = os.Args[2]
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	err = s.StopProcess(name)
	if err != nil {
		log.Fatal(err)
	}
}

func restart() {
	name := ""
	if len(os.Args) == 3 {
		name = os.Args[2]
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	err = s.RestartProcess(name)
	if err != nil {
		log.Fatal(err)
	}
}

func env() {
	name := ""
	if len(os.Args) == 3 {
		name = os.Args[2]
	} else {
		log.Fatal("Missing name")
		usage()
	}

	s, err := screend.InitScreend()
	if err != nil {
		log.Fatal(err)
	}
	p, err := s.ListProcesses(false, name)
	if err != nil {
		log.Fatal(err)
	}
	if len(p) == 0 {
		log.Fatal("Process not found")
	}
	for k, v := range p[0].Env {
		println(k + "=" + v)
	}
}

func setEnv() {
	log.Fatal("Not implemented")
}
