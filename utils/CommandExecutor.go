package utils

import (
	"bytes"
	"fmt"
	"github.com/devtron-labs/common-lib/utils/secretScanner"
	"io"
	"os"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd) error {
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)
	cmd.Stdout = mw
	cmd.Stderr = mw
	if err := cmd.Run(); err != nil {
		return err
	}
	//log.Println(stdBuffer.String())
	return nil
}

func RunCommandWithSecretMasking(cmd *exec.Cmd) error {

	var outBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		fmt.Printf("Command execution failed: %v\n", err)
		os.Exit(1)
		return err
	}

	buf := new(bytes.Buffer)
	// Call the function to mask secrets and print the masked output
	maskedStream, err := secretScanner.MaskSecretsStream(&outBuf)
	if err != nil {
		fmt.Printf("Error masking secrets: %v\n", err)
		os.Exit(1)
		return err
	}
	_, er := buf.ReadFrom(maskedStream)
	if er != nil {
		fmt.Printf("Error reading from masked stream: %v\n", er)
		os.Exit(1)
		return err
	}
	fmt.Println(buf.String())
	return nil
}
