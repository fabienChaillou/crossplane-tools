package core

import "os"

func read(file string) ([]byte, error) {
	return os.ReadFile(file)
}
