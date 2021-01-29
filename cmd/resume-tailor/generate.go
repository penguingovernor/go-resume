package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/penguingovernor/go-resume/internal/resume"
	"gopkg.in/yaml.v2"
)

func generatePDFResume(res *resume.Resume, url string, dest io.Writer) error {
	// Marshal the JSON.
	var jsonBody bytes.Buffer
	err := json.NewEncoder(&jsonBody).Encode(res)
	if err != nil {
		return err
	}

	// Send it off.
	contentType := "application/json"
	resp, err := http.Post(url, contentType, &jsonBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != http.StatusOK {
		return fmt.Errorf("pdf generation with error: %d %s", status, http.StatusText(status))
	}

	// Copy the message to the writer.
	_, err = io.Copy(dest, resp.Body)
	return err
}

func mergeSkillsAndResume(yamlSkillsRD, jsonResumeRD io.Reader) (*resume.Resume, error) {
	// Parse it.
	var skills []resume.Skills
	if err := yaml.NewDecoder(yamlSkillsRD).Decode(&skills); err != nil {
		return nil, err
	}

	// Now parse the base json resume
	var jsonResume resume.Resume
	if err := json.NewDecoder(jsonResumeRD).Decode(&jsonResume); err != nil {
		return nil, err
	}

	// "Merge" them together.
	jsonResume.Skills = skills

	return &jsonResume, nil
}
