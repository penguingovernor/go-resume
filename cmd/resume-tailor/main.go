package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	pdf "github.com/pdfcpu/pdfcpu/pkg/api"
)

func mainError() error {
	// Get the flags.
	flags, err := getFlags()
	if err != nil {
		return err
	}

	// Open the skills file.
	yamlFD, err := os.Open(flags.skillsFileName)
	if err != nil {
		return err
	}
	defer yamlFD.Close()

	// Open the JSON file.
	jsonFD, err := os.Open(flags.jsonResumeFileName)
	if err != nil {
		return err
	}
	defer jsonFD.Close()

	// Merge the two files, to create a customized resume json.
	resume, err := mergeSkillsAndResume(yamlFD, jsonFD)
	if err != nil {
		return err
	}

	// Set the template number.
	resume.SelectedTemplate = flags.templateNo

	// Now output the file.

	// In JSON, if specified.
	if flags.JSON {
		return json.NewEncoder(os.Stdout).Encode(resume)
	}

	// If we're forcing a single page.
	if flags.forceSingle {
		// Then create the pdf in a buffer.
		var buffer bytes.Buffer
		if err := generatePDFResume(resume, flags.url, &buffer); err != nil {
			return err
		}
		// Trim and rewrite it to stdout.
		return pdf.RemovePages(bytes.NewReader(buffer.Bytes()), os.Stdout, []string{"2-"}, nil)
	}

	// Otherwise just generate it to stdout.
	return generatePDFResume(resume, flags.url, os.Stdout)
}

func main() {
	if err := mainError(); err != nil {
		log.Fatalf("Application failed with error: %s\n", err)
	}
}
