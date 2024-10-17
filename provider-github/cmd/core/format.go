package core

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	pattern = "%v * time.Minute"
	match   = "o.PollInterval"
)

func init() {
	for i, config := range DataConfig.Config {
		fmt.Println(i, config)
		format(config.File, config.Pattern)
	}
}

func format(file, time string) {
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
		// log.Println("Line fille: ", line)
		if strings.Contains(line, "o.PollInterval") {
			log.Println("Line fille: ", line)
			// new := strings.Replace(line, match, pattern, -1)
			// p, _ := fmt.Sprintf(pattern, time)
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
