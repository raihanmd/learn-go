package utils

import (
	"flag"
	"fmt"
)

// untuk CLI an -foo -bar

func Flag() {
	host := flag.String("host", "localhost", "host")
	username := flag.String("username", "root", "username")
	password := flag.String("password", "", "password")

	flag.Parse()

	fmt.Println("========== Reading Flag ==========")
	fmt.Printf("Host Flag: %v\n", *host)
	fmt.Printf("Username Flag: %v\n", *username)
	fmt.Printf("Password Flag: %v\n\n\n", *password)
}
