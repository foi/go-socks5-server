# go socks5 server

Because it is primitive but modern.

Authorization via config (/etc/go-socks5-server.config.json), not via linux PAM.

Socks5 proxy server written on Go (based on: https://github.com/armon/go-socks5)

## how to install

### CentOS6 (tested) or another sysvinit legacy OSes (not tested)

```
cd /tmp
curl -L -O https://github.com/foi/go-socks5-server/releases/download/1.1/go-socks5-server-1.1.0-1.x86_64.legacy.rpm
sudo yum install go-socks5-server-1.1.0-1.x86_64.legacy.rpm
sudo chkconfig --add go-socks5-server
sudo chkconfig go-socks5-server on
sudo service go-socks5-server start
# open port in firewall
```

### Debian 9+/Ubuntu 16.04+

```
cd /tmp
wget https://github.com/foi/go-socks5-server/releases/download/1.1/go-socks5-server_1.1.0-1_amd64.deb
sudo dpkg -i go-socks5-server_1.1.0-1_amd64.deb
# edit config
# sudo nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl start go-socks5-server
sudo systemctl enable go-socks5-server

```
### Centos 7/Fedora

```
cd /tmp
curl -L -O https://github.com/foi/go-socks5-server/releases/download/1.1/go-socks5-server-1.1.0-1.x86_64.rpm
yum install -y go-socks5-server-1.1.0-1.x86_64.rpm
# edit config
# sudo nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl start go-socks5-server
sudo systemctl enable go-socks5-server

```

### Arch linux/Manjaro

```
yaourt go-socks5-server-bin
# edit config
# sudo nano /etc/go-socks5-server.config.json
sudo systemctl daemon-reload
sudo systemctl restart go-socks5-server
sudo systemctl enable go-socks5-server

```

### Do not forget to change the config in /etc/go-socks5-server.config.json and restart service
