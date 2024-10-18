/*
Copyright 2021 Upbound Inc.
*/

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/crossplane/upjet/pkg/pipeline"

	loader "github.com/upbound/upjet-provider-template/cmd/core"
	"github.com/upbound/upjet-provider-template/config"
)

func main() {
	if len(os.Args) < 2 || os.Args[1] == "" {
		panic("root directory is required to be given as argument")
	}

	rootDir := os.Args[1]
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}

	log.Println("Dir is: ", absRootDir)

	pipeline.Run(config.GetProvider(), absRootDir)

	loader.Run()
}
