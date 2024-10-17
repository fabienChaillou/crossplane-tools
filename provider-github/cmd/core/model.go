package core

type Config struct {
	Config []struct {
		Name    string `yaml:"name"`
		File    string `yaml:"file"`
		Path    string `yaml:"path"`
		Pattern string `yaml:"pattern"`
		Match   string `yaml:"match"`
	} `yaml:"config"`
}
