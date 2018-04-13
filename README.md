# go socks5 server

Because it is primitive but modern

Socks5 proxy server written on Go (based on: https://github.com/armon/go-socks5)

## how to install

### Debian/Ubuntu

```
cd /tmp
wget https://github.com/foi/go-socks5-server/releases/download/1.0/go-socks5-server_1.0.0-1_amd64.deb
sudo dpkg -i go-socks5-server_1.0.0-1_amd64.deb
# edit config
# nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl start go-socks5-server
sudo systemctl enable go-socks5-server

```
### Centos/Fedora

```
cd /tmp
curl -L -O https://github.com/foi/go-socks5-server/releases/download/1.0/go-socks5-server-1.0.0-1.x86_64.rpm
yum install -y go-socks5-server-1.0.0-1.x86_64.rpm
# edit config
# nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl start go-socks5-server
sudo systemctl enable go-socks5-server

```

### Do not forget to change the config in /etc/go-socks5-server.config.json and restart service
