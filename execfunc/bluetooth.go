package execfunc

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/keremimo/initialize-arch/credmanagement"
)

func EnableBluetooth(c *credmanagement.Credentials) error {
	cmd := exec.Command("sudo", "-S", "systemctl", "enable", "--now", "bluetooth")

	cmd.Stdin = bytes.NewBufferString(c.Password + "\n")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error enabling Bluetooth: %v", err)
	}

	fmt.Println("Bluetooth service enabled and started successfully.")
	return nil
}
