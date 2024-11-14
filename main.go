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

	return nil
}

func main() {
	credentials := new(Credentials)
	err := InitializeCredentials(credentials)
	if err != nil {
		fmt.Println("Something horrible happened.")
		fmt.Println(err)
	}
	fmt.Printf("Your username is %s, your email is %s and your password is safe with us.", credentials.Username, credentials.Email)
}
