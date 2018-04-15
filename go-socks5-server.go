package main

import socks5 "github.com/armon/go-socks5"
import (
	"encoding/json"
	"fmt"
	"os"
)

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
	file, _ := os.Open("/etc/go-socks5-server.config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}

	creds := socks5.StaticCredentials{}

	for _, c := range configuration.Credentials {
		creds[c.User] = c.Pass
	}

	cator := socks5.UserPassAuthenticator{Credentials: creds}

	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
	}

	server, err := socks5.New(conf)

	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("go-socks5-server started successfully: ip %s port %s", configuration.Ip, configuration.Port))

	if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%s", configuration.Ip, configuration.Port)); err != nil {
		panic(err)
	}
}
