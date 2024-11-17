package credmanagement

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type Credentials struct {
	Username    string
	Password    string
	Email       string
	GithubName  string
	GithubToken string
}

func InitializeCredentials(c *Credentials) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	c.Username = strings.TrimSpace(username)
	fmt.Println("Enter your e-mail address: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	c.Email = strings.TrimSpace(email)
	fmt.Println("Input your GitHub name: ")
	githubName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	c.GithubName = strings.TrimSpace(githubName)
	fmt.Println("Input a password")
	password, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err
	}
	c.Password = string(password)

	return nil
}
