package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
)

const (
	exitCodeSuccess = iota
	exitCodeFailure
)

// actualMain is used here to ensure all deferred calls are executed.
func actualMain(flags cliFlags) int {
	// Open the log file.
	logFile, err := os.OpenFile(flags.logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer logFile.Close()
	if err != nil {
		log.Printf("unable to open log file: %v\n", err)
		return exitCodeFailure
	}

	// Create a logger from it.
	logger := log.New(logFile, "", log.LstdFlags)

	// Save the working directory.
	pwd, err := os.Getwd()
	defer os.Chdir(pwd)
	if err != nil {
		logger.Printf("unable to get workding directory: %v\n", err)
		return exitCodeFailure
	}

	// Change to the resumake.io directory.
	if err := os.Chdir(flags.resumeGitDirectory); err != nil {
		logger.Printf("unable to change into the resumake.io direcotry: %v\n", err)
		return exitCodeFailure
	}

	// Run the build script.
	if !flags.skipBuild {
		npmBuild := exec.Command("npm", "run", "build")
		npmBuild.Stdout = logFile
		npmBuild.Stderr = logFile

		fmt.Println("Building resumake.io, this may take some time...")

		if err := npmBuild.Run(); err != nil {
			logger.Printf("failed to build resumake.io: %v\n", err)
			return exitCodeFailure
		}
	}

	// Configure npm start.
	npmStart := exec.Command("npm", "start")
	npmStart.Stderr = logFile
	npmStart.Stdout = logFile

	// Start the command.
	if err := npmStart.Start(); err != nil {
		logger.Printf("failed to start npm start: %v\n", err)
		return exitCodeFailure
	}

	// The ports come from here:
	// https://github.com/saadq/resumake.io/blob/master/contributing.md#project-overview
	fmt.Println("Client started at http://localhost:3000")
	fmt.Println("Server started at http://localhost:3001")

	// Setup a trap routine.
	sigChan := make(chan os.Signal)
	doneChan := make(chan error)
	signal.Notify(sigChan, os.Interrupt, os.Kill)

	// Wait for the command to finish in the background.
	go func() {
		doneChan <- npmStart.Wait()
	}()

	// Run until one of the channels sends a value.
outerFor:
	for {
		select {

		// If we can read from the trap.
		case <-sigChan:
			{
				// Get the pid incase we can't kill the process.
				id := npmStart.Process.Pid
				if err := npmStart.Process.Kill(); err != nil {
					logger.Printf("failed to killed process (PID: %d): %v\n", id, err)
					return exitCodeFailure
				}

				// Then continue as normal.
				break outerFor
			}

		// If the command finished.
		case err := <-doneChan:
			{
				// Check for any errors and report them.
				if err != nil {
					logger.Printf("npm start failed to finish: %v\n", err)
					return exitCodeFailure
				}
				// Otherwise continue as normal.
				break outerFor
			}
		}
	}

	return exitCodeSuccess
}

func main() {
	// Get the flags.
	flags := getFlags()

	// Run the application.
	exitCode := actualMain(flags)

	// Check for errors.
	if exitCode != exitCodeSuccess {
		log.Printf("An error occurred. Please check %s for more information\n", flags.logFile)
	}

	// Then exit.
	os.Exit(exitCode)
}
