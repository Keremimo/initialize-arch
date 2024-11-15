package execfunc

import (
	"fmt"
	"github.com/keremimo/initialize-arch/credmanagement"
	"os/exec"
)

func EnableBluetooth(c *credmanagement.Credentials) error {
	out, err := exec.Command("/bin/sh", "-c", "echo '%s' | sudo -S systemctl enable --now bluetooth", c.Password).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out[:]))
	return nil
}
