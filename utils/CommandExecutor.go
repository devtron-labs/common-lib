package utils

import (
	"fmt"
	"github.com/devtron-labs/common-lib/utils/secretScanner"
	"io"
	"os"
	"os/exec"
)

var maskSecrets = true

func RunCommand(cmd *exec.Cmd) error {
	// Create a pipe for the command's stdout
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error creating stdout pipe: %v\n", err)
		return err
	}
	cmd.Stderr = cmd.Stdout // Combine stderr and stdout

	// Start the command
	if err := cmd.Start(); err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
		return err
	}

	if maskSecrets {
		// Create a goroutine to handle real-time output processing
		go func() {
			// Wrap the pipe reader to mask secrets
			maskedStream, err := secretScanner.MaskSecretsStream(stdoutPipe)
			if err != nil {
				fmt.Printf("Error masking secrets: %v\n", err)
				return
			}

			// Copy the masked stream to stdout
			if _, err := io.Copy(os.Stdout, maskedStream); err != nil {
				fmt.Printf("Error reading masked stream: %v\n", err)
				return
			}
		}()
	} else {
		// Copy the output directly to stdout
		if _, err := io.Copy(os.Stdout, stdoutPipe); err != nil {
			fmt.Printf("Error reading stream: %v\n", err)
			return err
		}
	}

	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
		return err
	}

	return nil
}
