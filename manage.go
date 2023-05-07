package screend

import (
	"errors"
)

func (s *Screend) AddProcess(p Process, force bool) (err error) {

	// check if process exists
	for _, process := range s.config.Processes {
		if process.Name == p.Name {
			if force {
				err = s.RemoveProcess(p.Name)
				if err != nil {
					return
				}
			} else {
				err = errors.New("process already exists")
				return
			}
		}
	}

	// add process
	s.config.Processes = append(s.config.Processes, p)

	// write config file
	err = writeConfig(s.config)

	return
}

func (s *Screend) RemoveProcess(name string) (err error) {
	found := name == "--all"
	// check if process exists
	for i, process := range s.config.Processes {
		if process.Name == name || name == "--all" {
			// remove process

			s.config.Processes = append(s.config.Processes[:i], s.config.Processes[i+1:]...)
			err = writeConfig(s.config)
			found = true
		}
	}

	if !found {
		err = errors.New("process does not exist")
	}
	return
}
