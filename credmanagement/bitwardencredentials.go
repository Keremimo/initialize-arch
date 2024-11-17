package credmanagement

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

type BwCredentials struct {
	Server    string
	Username  string
	Password  string
	TwoFactor string
	Session   string
}

func CreateBitwardenAuth(c *BwCredentials) error {
	fmt.Println("Input your 2FA code: ")
	twoFactor, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err
	}
	c.TwoFactor = string(twoFactor)

	return nil
}
