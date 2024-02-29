package main

import (
	"flag"
	"fmt"
)

func main() {

	var host *string = flag.String("host", "localhost", "Host database")
	var port *int = flag.Int("port", 0, "Port database")
	var username *string = flag.String("username", "root", "Username database")
	var password *string = flag.String("password", "root", "Password database")

	flag.Parse()

	fmt.Printf("Host : %s\n", *host)
	fmt.Printf("Port : %d\n", *port)
	fmt.Printf("Username : %s\n", *username)
	fmt.Printf("Password : %s\n", *password)

}
