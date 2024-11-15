package main

import (
	"bytes"
	"fmt"
	"github.com/keremimo/initialize-arch/credmanagement"
	"os/exec"
)

func InstallGithubCli(c *credmanagement.Credentials) error {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo '%s' | sudo -S pacman -S --noconfirm github-cli bitwarden-cli", c.Password))

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

func EnableBluetooth(c *credmanagement.Credentials) error {
	out, err := exec.Command("/bin/sh", "-c", "echo '%s' | sudo -S systemctl enable --now bluetooth", c.Password).Output()
	if err != nil {
		return err
	}
	fmt.Println(string(out[:]))
	return nil
}

func main() {
	credentials := new(credmanagement.Credentials)
	bwCredentials := new(credmanagement.BwCredentials)
	_ = bwCredentials // Placeholder
	err := credmanagement.InitializeCredentials(credentials)
	if err != nil {
		fmt.Println("Something horrible happened.")
		fmt.Println(err)
	}
	// err = EnableBluetooth(credentials)
	err = InstallGithubCli(credentials)
	if err != nil {
		fmt.Println(err)
	}
}
