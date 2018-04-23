package main

import (
	"encoding/json"
	"fmt"
	"os"
	"net"
	"flag"
	"github.com/genevieve/go-socks5"
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

	fmt.Println(fmt.Sprintf("Starting server on: %s:%s",configuration.Ip, configuration.Port));

	if err := server.ListenAndServe("tcp", net.JoinHostPort(configuration.Ip, configuration.Port));
  err != nil {
		panic(err)
	}
}
