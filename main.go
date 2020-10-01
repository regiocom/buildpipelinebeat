package main

import (
	"os"

	"github.com/regiocom/buildpipelinebeat/cmd"

	_ "github.com/regiocom/buildpipelinebeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
