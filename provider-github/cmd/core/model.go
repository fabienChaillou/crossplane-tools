package core

type Config struct {
	Config []struct {
		Name string `yaml:"name"`
		File string `yaml:"file"`
		Path string `yaml:"path"`
		Time string `yaml:"time"`
	} `yaml:"config"`
}
