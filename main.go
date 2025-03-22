package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/pragus/go-socks5"
)

const config = "/etc/go-socks5-server.config.json"

var Version string

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
	path := flag.String(
		"config",
		config,
		"config file path",
	)
	flag.Parse()

	file, err := os.Open(*path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

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

	fmt.Printf(
		"Starting go-socks5-server (%s) on: %s:%s",
		Version,
		configuration.Ip,
		configuration.Port,
	)

	if err := server.ListenAndServe("tcp", net.JoinHostPort(configuration.Ip, configuration.Port)); err != nil {
		panic(err)
	}
}
