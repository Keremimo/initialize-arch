package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/keremimo/initialize-arch/credmanagement"
	"github.com/keremimo/initialize-arch/execfunc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	credentials := &credmanagement.Credentials{
		Username:   os.Getenv("USERNAME"),
		Password:   os.Getenv("PASSWORD"),
		Email:      os.Getenv("EMAIL"),
		GithubName: os.Getenv("GITHUB_NAME"),
	}

	bwCredentials := &credmanagement.BwCredentials{
		Server:   os.Getenv("BW_SERVER"),
		Username: os.Getenv("BW_USERNAME"),
		Password: os.Getenv("BW_PASSWORD"),
	}
	_ = bwCredentials // Placeholder
	// err = credmanagement.InitializeCredentials(credentials)
	// if err != nil {
	// 	fmt.Println("Something horrible happened.")
	// 	fmt.Println(err)
	// }
	err = execfunc.EnableBluetooth(credentials)
	if err != nil {
		fmt.Println(err)
	}
	err = execfunc.InstallPackages(credentials, "bitwarden-cli github-cli")
	if err != nil {
		fmt.Println(err)
	}

}
