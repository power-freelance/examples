#!/usr/bin/env bash

export DEBIAN_FRONTEND=noninteractive

# Post install steps for ubuntu server.
apt update -yq
apt upgrade -yq

adduser --disabled-password --gecos "" ubuntu
usermod -aG sudo ubuntu
echo "ubuntu ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

sed -i "s/#Port 22/Port 2222/" /etc/ssh/sshd_config
service sshd restart

ufw allow 2222
ufw allow 80
ufw allow 443
ufw --force enable

rsync --archive --chown=ubuntu:ubuntu ~/.ssh /home/ubuntu

# Install docker

# Requirements
apt update
apt install python-minimal apt-transport-https ca-certificates curl gnupg-agent software-properties-common -yq

# Docker
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt update -yq
apt install docker-ce docker-ce-cli containerd.io -yq
usermod -aG docker ubuntu
systemctl enable docker
systemctl start docker

# Docker compose
curl -L "https://github.com/docker/compose/releases/download/1.24.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose
