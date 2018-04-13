package main

import socks5 "github.com/armon/go-socks5"
import "fmt"
import "flag"

func main(){
  ip := flag.String("ip", "0.0.0.0", "")
  port := flag.String("port", "1080", "")
  user := flag.String("user", "changeme", "username")
  pass := flag.String("pass", "andme", "password")

  flag.Parse()

  creds := socks5.StaticCredentials{
    *user: *pass,
  }

  cator := socks5.UserPassAuthenticator{Credentials: creds}

  conf := &socks5.Config{
    AuthMethods: []socks5.Authenticator{cator},
  }

  server, err := socks5.New(conf)

  fmt.Println("go-socks5-server started successfully.")

  if err != nil {
    panic(err)
  }

  if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%s", *ip, *port));
  err != nil {
    panic(err)
  }
}
