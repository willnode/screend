package screend

type Screend struct {
	config Config
}

func InitScreend() (instance *Screend, err error) {
	instance = &Screend{}
	err = instance.ReadConfig()
	return
}

func (s *Screend) ReadConfig() (err error) {
	s.config, err = readConfig()
	return
}
