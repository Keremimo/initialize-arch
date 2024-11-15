package execfunc

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/keremimo/initialize-arch/credmanagement"
)

func InstallPackages(c *credmanagement.Credentials, p string) error {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo '%s' | sudo -S pacman -S --noconfirm --needed %s", c.Password, p))

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
