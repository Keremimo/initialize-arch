package execfunc

import (
	"fmt"
	"github.com/keremimo/initialize-arch/credmanagement"
	"os/exec"
)

func BitwardenLogin(c *credmanagement.Credentials, b *credmanagement.BwCredentials) error {
	out, err := exec.Command("/bin/sh", "-c", "echo %s | bw config server", b.Server).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out[:]))

	out, err = exec.Command("/bin/sh", "-c", fmt.Sprintf("echo %s %s --code %s | bw login", b.Username, b.Password, b.TwoFactor)).Output()

	if err != nil {
		return err
	}
	fmt.Println(string(out[:]))

	return nil
}
