package main

import (
	"os"

	"github.com/renato0307/outsystemslogsbeat/cmd"

	_ "github.com/renato0307/outsystemslogsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
