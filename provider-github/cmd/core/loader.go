package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	pattern = "%v * time.Minute"
	match   = "o.PollInterval"
)

var DataConfig Config

func init() {
	loadConfigs()
}

func Run() {
	e := os.Getenv("FORMAT")

	if len(e) == 0 {
		log.Fatalln("Don't parse the formate controller file!")
		return
	}

	// ParseBool returns the boolean value represented by the string.
	// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
	t, err := strconv.ParseBool(e)
	if err != nil {
		log.Fatalln(err)
		return
	}

	if !t {
		log.Println("Don't replace ", match, "into controller because Parse Environment is: ", t)
		return
	}

	for i, config := range DataConfig.Config {
		log.Println(i, config)
		replace(config.Path+config.File, config.Time)
	}
}

func loadConfigs() {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	data, err := read(dir + "/cmd/config/config.yaml")

	if err != nil {
		log.Fatal("Error while reading app config file", err)
	}

	laodServerConfig(data)

}

func laodServerConfig(yamlData []byte) {
	yaml.Unmarshal(yamlData, &DataConfig)
	fmt.Println("Loaded Server Config with ", len(DataConfig.Config), " config files")
}

func replace(file, time string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	f := dir + file

	input, err := read(f)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, match) {
			log.Println("Line fille: ", line)
			lines[i] = strings.Replace(line, match, fmt.Sprintf(pattern, time), -1)
		}
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(f, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("go", "fmt", f)
	stdout, err := cmd.Output()

	if err != nil {
		log.Fatalln(err)
		return
	}

	// Print the output
	log.Println(string(stdout))
}

func read(file string) ([]byte, error) {
	return os.ReadFile(file)
}
