package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
)

// actualMain is used here to ensure all deferred calls are executed.
func actualMain(flags cliFlags) int {
	const (
		successCode = iota
		errCode
	)

	// Open the log file.
	logFile, err := os.OpenFile(flags.logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("unable to open log file: %v\n", err)

	}

	// Change to the resumake.io directory.
	pwd, err := os.Getwd()
	defer os.Chdir(pwd)
	if err != nil {
		log.Printf("unable to get workding directory: %v\n", err)
		return errCode
	}

	if err := os.Chdir(flags.resumeGitDirectory); err != nil {
		log.Printf("unable to change into the resumake.io direcotry: %v\n", err)
		return errCode
	}

	// Run the build script.
	if !flags.skipBuild {
		npmBuild := exec.Command("npm", "run", "build")
		npmBuild.Stdout = logFile
		npmBuild.Stderr = logFile

		log.Println("Building resumake.io, this may take some time...")

		if err := npmBuild.Run(); err != nil {
			log.Printf("failed to build resumake.io: %v\n", err)
			log.Printf("check the log file at :%v\n", flags.logFile)
			return errCode
		}
	}

	// Configure npm start.
	npmStart := exec.Command("npm", "start")
	npmStart.Stderr = logFile
	npmStart.Stdout = logFile

	// Start the command.
	if err := npmStart.Start(); err != nil {
		log.Printf("failed to start npm start: %v\n", err)
		return errCode
	}

	// The ports come from here:
	// https://github.com/saadq/resumake.io/blob/master/contributing.md#project-overview
	log.Println("Client started at http://localhost:3000")
	log.Println("Server started at http://localhost:3001")

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
					log.Printf("failed to killed process (PID: %d): %v\n", id, err)
					return errCode
				}

				// Then continue as normal.
				break outerFor
			}

		// If the command finished.
		case err := <-doneChan:
			{
				// Check for any errors and report them.
				if err != nil {
					log.Printf("npm start failed to finish: %v\n", err)
					return errCode
				}
				// Otherwise continue as normal.
				break outerFor
			}
		}
	}

	return successCode
}

func main() {
	// Get the flags.
	flags := getFlags()

	// Run the application.
	exitCode := actualMain(flags)

	os.Exit(exitCode)
}
