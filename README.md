# go socks5 server

Primitive but modern.

Authorization via config (/etc/go-socks5-server.config.json), not via linux PAM.

Socks5 proxy server written on Go (based on: https://github.com/armon/go-socks5)

## how to install

### Debian 9+/Ubuntu 16.04+

```
cd /tmp
wget https://github.com/foi/go-socks5-server/releases/download/v1.1.0/go-socks5-server_1.1.0-1_amd64.deb
sudo dpkg -i go-socks5-server_1.1.0-1_amd64.deb
# edit config
# sudo nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl start go-socks5-server
sudo systemctl enable go-socks5-server
# open port in firewall
sudo ufw allow 22
sudo ufw allow YOUR_CONFIGURED_GO_SOCKS5_SERVER_PORT
sudo ufw disable && sudo ufw enable
```
### Centos 7/Fedora

```
cd /tmp
curl -L -O https://github.com/foi/go-socks5-server/releases/download/v1.1.0/go-socks5-server-1.1.0-1.x86_64.rpm
yum install -y go-socks5-server-1.1.0-1.x86_64.rpm
# edit config
# sudo nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl start go-socks5-server
sudo systemctl enable go-socks5-server
# open port in firewall
sudo firewall-cmd --permanent --add-port=YOUR_CONFIGURED_GO_SOCKS5_SERVER_PORT/tcp
sudo firewall-cmd --reload
```

### Arch linux/Manjaro

```
yaourt go-socks5-server-bin
# edit config
# sudo nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl restart go-socks5-server
sudo systemctl enable go-socks5-server
# open port in firewall
```

### How to check go socks5 server

`curl -x socks5://username:password@address:1080 https://web.telegram.org`

### Do not forget to change the config in /etc/go-socks5-server.config.json and restart service
