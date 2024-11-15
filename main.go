package main

import (
	"fmt"
	"github.com/keremimo/initialize-arch/credmanagement"
	"github.com/keremimo/initialize-arch/execfunc"
)

func main() {
	credentials := new(credmanagement.Credentials)
	bwCredentials := new(credmanagement.BwCredentials)
	_ = bwCredentials // Placeholder
	err := credmanagement.InitializeCredentials(credentials)
	if err != nil {
		fmt.Println("Something horrible happened.")
		fmt.Println(err)
	}
	err = execfunc.EnableBluetooth(credentials)
	if err != nil {
		fmt.Println(err)
	}
	err = execfunc.InstallPackages(credentials, "bitwarden-cli github-cli")
	if err != nil {
		fmt.Println(err)
	}

}
