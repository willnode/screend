package screend

import (
	"os"
	"os/exec"
	"regexp"
)

func (s *Screend) ListProcesses(withOsProcess bool, name string) (processes []Process, err error) {

	// get processes
	for _, process := range s.config.Processes {
		p := process
		if name != "" && p.Name != name {
			continue
		}
		processes = append(processes, p)
	}

	if withOsProcess {
		var output []byte
		cmd := exec.Command(s.config.Bin.Screen, "-ls")
		output, err = cmd.CombinedOutput()
		if err != nil {
			if cmd.ProcessState.ExitCode() != 1 {
				return
			}
			err = nil
		}
		// parse output
		r := regexp.MustCompile(`\t(\d+)\.([^.]+?)\t`)
		for _, match := range r.FindAllStringSubmatch(string(output), -1) {
			for i, process := range processes {
				if process.Name == match[2] {
					processes[i].ScreenID = match[1]
				}
			}
		}
	}

	return
}

func (s *Screend) StartProcess(name string) (err error) {
	procList, err := s.ListProcesses(true, name)
	if err != nil {
		return
	}

	for _, proc := range procList {
		if proc.ScreenID != "" {
			continue
		}

		args := []string{"-dmS", proc.Name, proc.Command}
		args = append(args, proc.Args...)

		cmd := exec.Command(s.config.Bin.Screen, args...)
		cmd.Dir = proc.Dir
		cmd.Env = os.Environ()
		for key, value := range proc.Env {
			cmd.Env = append(cmd.Env, key+"="+value)
		}
		err = cmd.Run()
		if err != nil {
			return
		}
	}

	return
}

func (s *Screend) StopProcess(name string) (err error) {
	procList, err := s.ListProcesses(true, name)
	if err != nil {
		return
	}

	for _, proc := range procList {
		if proc.ScreenID == "" {
			continue
		}

		err = exec.Command(s.config.Bin.Screen, "-XS", proc.ScreenID, "quit").Run()
		if err != nil {
			return
		}
	}

	return
}

func (s *Screend) RestartProcess(name string) (err error) {
	err = s.StopProcess(name)
	if err != nil {
		return
	}
	err = s.StartProcess(name)
	if err != nil {
		return
	}
	return
}
