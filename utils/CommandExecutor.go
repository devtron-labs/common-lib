package utils

import (
	"bytes"
	"fmt"
	"github.com/devtron-labs/common-lib/utils/secretScanner"
	"io"
	"os/exec"
)

var maskSecrets = true

func RunCommand(cmd *exec.Cmd) error {

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
		return err
	}
	outBuf := bytes.NewBuffer(output)
	if maskSecrets {
		buf := new(bytes.Buffer)
		// Call the function to mask secrets and print the masked output
		maskedStream, err := secretScanner.MaskSecretsStream(outBuf)
		if err != nil {
			fmt.Printf("Error masking secrets: %v\n", err)
			return err
		}
		_, err = io.Copy(buf, maskedStream)
		if err != nil {
			fmt.Printf("Error reading from masked stream: %v\n", err)
			return err
		}
		fmt.Println(buf.String())
	} else {
		fmt.Println(outBuf.String())
	}
	return nil
}
