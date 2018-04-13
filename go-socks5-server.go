package main

import socks5 "github.com/armon/go-socks5"
import ( "fmt"
         "encoding/json"
         "os"
)

type Config struct {
    Ip string
    Port string
    User string
    Pass string
}

func main(){
  file, _ := os.Open("/etc/go-socks5-server.config.json")
  defer file.Close()
  fmt.Println(&file)
  decoder := json.NewDecoder(file)
  configuration := Config{}
  err := decoder.Decode(&configuration)
  if err != nil {
    panic(err)
  }

  creds := socks5.StaticCredentials{
    configuration.User: configuration.Pass,
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

  if err := server.ListenAndServe("tcp", fmt.Sprintf("%s:%s", configuration.Ip, configuration.Port));
  err != nil {
    panic(err)
  }
}
