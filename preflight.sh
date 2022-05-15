#!/bin/bash

# lmnl. preflight for debian sid

# the server hosting lxmxnxl.com is nothing fancy.
# everything is held in ram, the entire server has 1gb.

# debian sid is used, rolling release is easier

apt install curl wget apt-transport-https dirmngr -y 

echo -e "deb http://deb.debian.org/debian/ sid main contrib non-free" > /etc/apt/sources.list

apt update -y

apt dist-upgrade -y 

apt autoremove -y

apt install tmux golang git curl htop build-essential nginx -y

echo "Ready..." >> /var/www/html/index.html

# TODO - add entry to nginx.conf for reverse proxy from :8666 to :80