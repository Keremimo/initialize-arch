package execfunc

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/keremimo/initialize-arch/credmanagement"
)

func BitwardenLogin(c *credmanagement.Credentials, b *credmanagement.BwCredentials) error {
	cmd := exec.Command("/bin/sh", "-c", "echo %s | bw config server", b.Server)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error: %v, stderr: %s", err, stderr.String())
	}
	fmt.Println(stdout.String())

	cmd = exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s %s --code %s | bw login", b.Username, b.Password, b.TwoFactor))

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error: %v, stderr: %s", err, stderr.String())
	}
	fmt.Println(stdout.String())

	return nil
}
