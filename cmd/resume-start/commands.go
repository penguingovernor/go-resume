package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/penguingovernor/go-resume/internal/errgroup"
)

// commandFromString builds a command from a string like "foo -bar baz".
// It returns nil with the empty string.
func commandFromString(command string) *exec.Cmd {
	fields := strings.Fields(command)
	if len(fields) < 1 {
		return nil
	}
	return exec.Command(fields[0], fields[1:]...)
}

// runCommand builds a command with the provided logger.
func runCommand(commandToRun string, logger *log.Logger) error {
	// Make the command argument.
	cmd := commandFromString(commandToRun)
	if cmd == nil {
		return fmt.Errorf("could not create command from: %s", commandToRun)
	}

	// Get the pipes for stderr and stdout.
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// Start the command.
	if err := cmd.Start(); err != nil {
		return err
	}

	// Scan their combined outputs.
	scanner := bufio.NewScanner(io.MultiReader(stderr, stdout))
	for scanner.Scan() {
		logger.Print(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		logger.Printf("Failed to log output: %s\n", err)
		return err
	}

	// Wait until the command is done.
	return cmd.Wait()
}

// buildServer builds resumake.io's
// server. All its output will be written to the provided writer.
//
// buildServer assumes that the working directory contains resumake.io's root package.json.
func buildServer(logger *log.Logger) error {
	return runCommand("npm run build:server", logger)
}

// buildClient builds resumake.io's client. All its output is written to the provided writer.
//
// buildClient assumes that the working directory contains resumake.io's root package.json.
func buildClient(logger *log.Logger) error {
	return runCommand("npm run build:client", logger)
}

// concurentBuildClientSerer builds both the client and the server concurrently.
// If any errors occur the first one is returned.
func concurentBuildClientSerer(logger *log.Logger) error {
	wg := &errgroup.Group{}
	wg.Go(func() error { return buildClient(logger) })
	wg.Go(func() error { return buildServer(logger) })
	return wg.Wait()
}

// runClientServer runs the server and optionally the client.
// It logs all output to the provided logger.
func runClientServer(runClient bool, logger io.Writer) error {
	// Setup the signal trap.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	// These are our running commands.
	const server = "npm run start:server"
	const client = "npm run start:client"

	// Make the commands.
	serverCmd := commandFromString(server)
	serverCmd.Stderr = logger
	serverCmd.Stdout = logger

	clientCmd := commandFromString(client)
	clientCmd.Stderr = logger
	clientCmd.Stdout = logger

	// If the user wants to run the client, run it.
	if runClient {
		fmt.Println("Started client at http://localhost:3000")
		if err := clientCmd.Start(); err != nil {
			return err
		}
	}

	// Start the server.
	fmt.Println("Started server at http://localhost:3001")
	if err := serverCmd.Start(); err != nil {
		return err
	}

	// Ensure we can get results from waiting.
	clientChan := make(chan error)
	serverChan := make(chan error)

	// If a client, wait on the value.
	if runClient {
		go func() { clientChan <- clientCmd.Wait() }()
	}
	// Wait on the server.
	go func() { serverChan <- serverCmd.Wait() }()

	select {

	// If we got an interrupt.
	case <-sigChan:
		{
			// Kill the server.
			if err := serverCmd.Process.Kill(); err != nil {
				return fmt.Errorf("failed to kill process %d: %w", serverCmd.Process.Pid, err)
			}
			fmt.Println("Server closed")

			// Conditionally, kill the client.
			if runClient {
				if err := clientCmd.Process.Kill(); err != nil {
					return fmt.Errorf("failed to kill process %d: %w", clientCmd.Process.Pid, err)
				}
				fmt.Println("Client Closed")
			}
		}

	// If the client finished unexpectedely.
	case err := <-clientChan:
		{
			// Tell the user.
			if err != nil {
				fmt.Printf("Client quit unexpectedly: %s\n", err)
			} else {
				fmt.Printf("Client quit unexpectedly")
			}

			// Kill the server.
			if err := serverCmd.Process.Kill(); err != nil {
				return fmt.Errorf("failed to kill process %d: %w", serverCmd.Process.Pid, err)
			}
			fmt.Println("Server closed")

		}

	// If the server finished unexpectedely.
	case err := <-serverChan:
		{
			// Tell the user.
			if err != nil {
				fmt.Printf("Server quit unexpectedly: %s\n", err)
			} else {
				fmt.Printf("Server quit unexpectedly")
			}

			// Conditionally kill the client.
			if runClient {
				if err := clientCmd.Process.Kill(); err != nil {
					return fmt.Errorf("failed to kill process %d: %w", clientCmd.Process.Pid, err)
				}
				fmt.Println("Client Closed")
			}

		}

	}

	return nil
}
