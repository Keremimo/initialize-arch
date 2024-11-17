package githubssh

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FetchGithubPAT retrieves the GitHub personal access token from Bitwarden
func FetchGithubPAT(bitwardenSession string) (string, error) {
	// Command to fetch the GitHub PAT from Bitwarden
	cmd := exec.Command("bw", "get", "item", "github_pat")
	cmd.Env = append(os.Environ(), "BW_SESSION="+bitwardenSession)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error fetching GitHub PAT: %v, stderr: %s", err, stderr.String())
	}

	// Extract the `github_pat` field from the item JSON
	output := stdout.String()
	start := strings.Index(output, `"github_pat":"`)
	if start == -1 {
		return "", fmt.Errorf("github_pat not found in the Bitwarden item")
	}
	start += len(`"github_pat":"`)
	end := strings.Index(output[start:], `"`)
	if end == -1 {
		return "", fmt.Errorf("malformed GitHub PAT in Bitwarden item")
	}
	return output[start : start+end], nil
}

// GenerateSSHKey generates a new SSH key pair at the specified location
func GenerateSSHKey(email string, keyPath string) error {
	// Ensure the key directory exists
	dir := filepath.Dir(keyPath)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("error creating directory for SSH key: %v", err)
	}

	// Command to generate the SSH key
	cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-f", keyPath, "-N", "")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error generating SSH key: %v", err)
	}
	return nil
}

// AddSSHKeyToGithub uploads the SSH public key to GitHub using gh CLI
func AddSSHKeyToGithub(title, publicKeyPath, githubPAT string) error {
	// Read the public key content
	pubKey, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("error reading SSH public key: %v", err)
	}

	// Command to add the SSH key to GitHub
	cmd := exec.Command("gh", "ssh-key", "add", "-", "--title", title)
	cmd.Stdin = bytes.NewBuffer(pubKey) // Provide public key content via stdin
	cmd.Env = append(os.Environ(), "GITHUB_TOKEN="+githubPAT)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error adding SSH key to GitHub: %v", err)
	}
	return nil
}

// Main function to execute the process
func SetupGithubSSH(bitwardenSession, itemName, email, title, keyPath string) error {
	// Step 1: Fetch the GitHub PAT
	githubPAT, err := FetchGithubPAT(bitwardenSession, itemName)
	if err != nil {
		return err
	}

	// Step 2: Generate SSH Key
	if err := GenerateSSHKey(email, keyPath); err != nil {
		return err
	}

	// Step 3: Add SSH Key to GitHub
	publicKeyPath := keyPath + ".pub"
	if err := AddSSHKeyToGithub(title, publicKeyPath, githubPAT); err != nil {
		return err
	}

	fmt.Println("SSH key successfully added to GitHub!")
	return nil
}