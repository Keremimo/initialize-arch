package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type Credentials struct {
	Username   string
	Password   string
	Email      string
	GithubName string
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

func EnableBluetooth(c *Credentials) {
	out, _ := exec.Command("/bin/sh", "-c", "echo '%s' | sudo -S systemctl enable --now bluetooth", c.Password).Output()

	fmt.Println(string(out[:]))
}

func main() {
	credentials := new(Credentials)
	err := InitializeCredentials(credentials)
	if err != nil {
		fmt.Println("Something horrible happened.")
		fmt.Println(err)
	}
	fmt.Printf("Your username is %q, github name is %q, your email is %q and your password is safe with us.", credentials.Username, credentials.GithubName, credentials.Email)

	EnableBluetooth(credentials)
}
