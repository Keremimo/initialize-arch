package main

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"syscall"
)

type Credentials struct {
	Username string
	Password string
	Email    string
}

func InitializeCredentials(c *Credentials) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	c.Username = username
	fmt.Println("Enter your e-mail address: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	c.Email = email
	fmt.Println("Input a password")
	password, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return err
	}
	c.Password = string(password)
	// TODO: Gotta remove those printlines below after everything's done.
	fmt.Println(c.Username)
	fmt.Println(c.Email)
	fmt.Println(c.Password)
	return nil
}

func main() {
	credentials := new(Credentials)
	InitializeCredentials(credentials)
}
