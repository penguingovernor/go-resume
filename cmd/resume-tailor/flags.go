package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultTemplateNo      = 6
	deafultURL             = "http://localhost:3001/api/generate/resume"
	defaultForceSinglePage = false
	defaultOutputJSON      = false
	defaultResumeName      = "resume.json"
	defaultSkillName       = "skills.yaml"
	minArgs                = 2
)

type cliFlags struct {
	skillsFileName     string
	jsonResumeFileName string
	forceSingle        bool
	JSON               bool
	templateNo         int
	url                string
}

func getFlags() (*cliFlags, error) {
	// Create and define the flags.
	flags := &cliFlags{}

	flag.IntVar(&flags.templateNo, "template", defaultTemplateNo, "which template to use [values: 1-9]")
	flag.StringVar(&flags.url, "URL", deafultURL, "the API endpoint of resumake.io")
	flag.BoolVar(&flags.forceSingle, "force-single", defaultForceSinglePage, "force the resulting resume to be a single page")
	flag.BoolVar(&flags.JSON, "JSON", defaultOutputJSON, "output JSON instead of PDF data")

	// Define the usage.
	flag.Usage = func() {
		fmt.Printf("Usage: %s [OPTIONS] [resume.json skills.yaml] > resume.{pdf,json}\n\n", os.Args[0])
		fmt.Println("[OPTIONS]")
		flag.PrintDefaults()
	}

	// Parse.
	flag.Parse()

	// Get the positional arguments and validate them.
	if nArgs := flag.NArg(); nArgs == minArgs {
		flags.jsonResumeFileName = flag.Arg(0)
		flags.skillsFileName = flag.Arg(1)
	} else if nArgs == 0 {
		flags.jsonResumeFileName = defaultResumeName
		flags.skillsFileName = defaultSkillName
	} else {
		return nil, fmt.Errorf("wrong amount of positional arguments: got %d wanted %d or 0", nArgs, minArgs)
	}

	// Ensure we have a valid template number.
	if !(1 <= flags.templateNo && flags.templateNo <= 9) {
		return nil, fmt.Errorf("invalid template number: got %d wanted [1-9]", flags.templateNo)
	}

	// We can't force a single page and output JSON.
	if flags.forceSingle && flags.JSON {
		return nil, fmt.Errorf("cannot force single page on non pdf data")
	}

	return flags, nil
}
