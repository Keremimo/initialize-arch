package execfunc

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/Netflix/go-expect"
	"github.com/keremimo/initialize-arch/credmanagement"
)

func BitwardenInit(b *credmanagement.BwCredentials) error {
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
	return nil
}

func BitwardenLogin(c *credmanagement.Credentials, b *credmanagement.BwCredentials) error {
	// Login to Bitwarden
	// cmd := exec.Command("bw", "login", "--raw", "--code", b.TwoFactor)
	// var stdout, stderr bytes.Buffer
	// cmd.Stdin = strings.NewReader(fmt.Sprintf("%s \n %s \n", b.Username, b.Password))
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	console, err := expect.NewConsole(expect.WithDefaultTimeout(5 * time.Second))
	if err != nil {
		return fmt.Errorf("failed to create console: %v", err)
	}
	defer console.Close()

	// Login to Bitwarden
	cmd := exec.Command("bw", "login", "--raw")
	cmd.Stdin = console.Tty()
	cmd.Stdout = console.Tty()
	cmd.Stderr = console.Tty()

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %v", err)
	}

	// Handle input delays
	fmt.Println("Waiting for 'Email address:' prompt...")
	console.ExpectString("Email address:")
	console.SendLine(b.Username)

	fmt.Println("Waiting for 'Master password:' prompt...")
	console.ExpectString("Master password:")
	console.SendLine(b.Password)

	fmt.Println("Waiting for 'Two-step login code:' prompt...")
	console.ExpectString("Two-step login code:")
	console.SendLine(b.TwoFactor)

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("failed to wait for command: %v", err)
	}

	output, _ := console.ExpectEOF()
	sessionToken := strings.TrimSpace(output)

	if sessionToken == "" {
		return fmt.Errorf("failed to extract session token")
	}
	// Split session data by newlines and take the second line
	lines := strings.Split(string(sessionToken), "\n")
	if len(lines) > 1 {
		b.Session = strings.TrimSpace(lines[1])
	} else {
		b.Session = ""
	}
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
