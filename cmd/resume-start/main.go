package main

import (
	"fmt"
	"log"
	"os"
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

	// Build our dependencies in an anonymous function.
	err = func() error {
		// If the skip build flag is specified then don't build anything.
		if flags.skipBuild {
			return nil
		}

		// If we're not building the client then only build the server.
		if flags.noClient {
			fmt.Println("Building server...")
			return buildServer(logger)
		}

		// We must be building both the client and server.
		fmt.Println("Building client and server concurrently...")
		return concurentBuildClientSerer(logger)
	}()

	if err != nil {
		logger.Printf("build error: %s\n", err)
		return exitCodeFailure
	}

	if err := runClientServer(!flags.noClient, logFile); err != nil {
		logger.Printf("failed to start: %s\n", err)
		return exitCodeFailure
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
