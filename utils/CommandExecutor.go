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
	output, outputerr := cmd.CombinedOutput()

	outBuf := bytes.NewBuffer(output)
	if maskSecrets {
		buf := new(bytes.Buffer)
		// Call the function to mask secrets and print the masked output
		maskedStream, err := secretScanner.MaskSecretsStream(outBuf)
		if err != nil {
			fmt.Printf("Error masking secrets: %v\n", err)
			fmt.Println(outBuf.String())
		}
		_, err = io.Copy(buf, maskedStream)
		if err != nil {
			fmt.Printf("Error reading from masked stream: %v\n", err)
			fmt.Println(outBuf.String())
		}
		fmt.Println(buf.String())
	} else {
		fmt.Println(outBuf.String())
	}
	if outputerr != nil {
		fmt.Printf("Command execution failed: %v\n", outputerr)
		return outputerr
	}
	return nil
}
