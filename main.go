package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	doIt             = flag.Bool("DOIT", false, "actually delete; definitely destructive")
	deleteCount      = flag.Int("DELETE", 5, "delete how many per run?")
	domain           = flag.String("d", "", "domain")
	clientSecretFile = flag.String("c", "client_secret.json", "client_secret file")
	token            = flag.String("t", "", "token")
	approval         = flag.String("a", "", "approval")
)

func main() {
	flag.Parse()

	cs, err := getClientSecret(*clientSecretFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch {
	case token != nil && *token != "":
		if domain != nil && *domain == "" {
			fmt.Println("ERROR: Must provide domain.")
			flag.Usage()
			os.Exit(3)
		}
		err := enterTheIncinerator(cs, *token, *domain)
		if err != nil {
			fmt.Println(err)
			os.Exit(4)
		}
		return
	case approval != nil && *approval != "":
		token, err := getToken(cs, *approval)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Println("Token= ", token.AccessToken)
		fmt.Println("Expire= ", token.Expiry)
		return
	default:
		fmt.Println("Go to: ", getAuthCodeURL(cs))
		return
	}
}
