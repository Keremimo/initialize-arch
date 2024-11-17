package execfunc

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/keremimo/initialize-arch/credmanagement"
)

func BitwardenLogin(c *credmanagement.Credentials, b *credmanagement.BwCredentials) error {
	// Configure Bitwarden server
	cmd := exec.Command("bw", "config", "server", b.Server)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error: %v, stderr: %s", err, stderr.String())
	}
	fmt.Println(stdout.String())

	// Login to Bitwarden
	cmd = exec.Command("bw", "login", "--code", b.TwoFactor)
	cmd.Stdin = strings.NewReader(fmt.Sprintf("%s %s", b.Username, b.Password))
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error: %v, stderr: %s", err, stderr.String())
	}
	sessionOutput := stdout.String()
	fmt.Println(sessionOutput)

	// Extract the session token from the output
	sessionToken := extractSessionToken(sessionOutput)
	if sessionToken == "" {
		return fmt.Errorf("failed to extract session token")
	}
	b.Session = sessionToken

	return nil
}

// Helper function to extract the session token from the output
func extractSessionToken(output string) string {
	// Assuming the session token is printed in the output as "export BW_SESSION='<token>'"
	start := strings.Index(output, "export BW_SESSION='")
	if start == -1 {
		return ""
	}
	start += len("export BW_SESSION='")
	end := strings.Index(output[start:], "'")
	if end == -1 {
		return ""
	}
	return output[start : start+end]
}
