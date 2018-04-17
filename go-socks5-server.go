package main

import "github.com/armon/go-socks5"
import (
	"encoding/json"
	"fmt"
	"os"
	"flag"
)

const CFG = "/etc/go-socks5-server.config.json"

type Config struct {
	Ip          string
	Port        string
	Credentials []Credentials
}

type Credentials struct {
	User string
	Pass string
}

func main() {
	path := flag.String("config", CFG, "config file path")
	flag.Parse()

	file, err := os.Open(*path)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	configuration := Config{}
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	creds := socks5.StaticCredentials{}

	for _, c := range configuration.Credentials {
		creds[c.User] = c.Pass
	}

	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{
			socks5.UserPassAuthenticator{
				Credentials: creds,
			},
		},
	}

	server, err := socks5.New(conf)

	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp",
		fmt.Sprintf("%s:%s",
			configuration.Ip,
			configuration.Port)); err != nil {
		panic(err)

	} else {
		fmt.Println(
			fmt.Sprintf("go-socks5-server started successfully: ip %s port %s",
				configuration.Ip,
				configuration.Port))
	}
}
