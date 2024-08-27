package main

import (
	"github.com/Mzack9999/shuffledns/pkg/runner"
	"github.com/projectdiscovery/gologger"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()

	massdnsRunner, err := runner.New(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}

	massdnsRunner.RunEnumeration()
	massdnsRunner.Close()
}
