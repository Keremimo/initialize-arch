package execfunc

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/keremimo/initialize-arch/credmanagement"
)

func InstallPackages(c *credmanagement.Credentials, p string) error {
	packages := strings.Fields(p)
	if len(packages) == 0 {
		return fmt.Errorf("no packages specified")
	}

	// Command to update
	fmt.Println("Updating package database...")
	updateCmd := exec.Command("sudo", "pacman", "-Sy")
	updateCmd.Stdin = bytes.NewBufferString(c.Password + "\n")
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr

	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("error updating package database: %v", err)
	}

	args := append([]string{"pacman", "-S", "--noconfirm", "--needed"}, packages...)
	fmt.Printf("Installing packages: %v\n", packages)
	installCmd := exec.Command("sudo", args...)
	installCmd.Stdin = bytes.NewBufferString(c.Password + "\n")
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	// Execute the install command
	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("error installing packages: %v", err)
	}

	fmt.Println("Packages installed successfully.")
	return nil
}
