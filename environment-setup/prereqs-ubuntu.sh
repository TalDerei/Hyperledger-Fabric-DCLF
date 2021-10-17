#!/bin/bash
# Prereqsuisites referencing: https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html

# cURL
## Step 1. Update the package list
sudo apt-get update

## Step 2. Install cURL
sudo apt install curl

## Step 3. Check curl version
curl --version


# Docker
## Step 1. Update the package list
sudo apt-get update

## Step 2. Install the required packages
sudo apt-get install apt-transport-https ca-certificates gnupg-agent software-properties-common

## Step 3. Add Docker’s official GPG key:
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

## Step 4. Setup the Docker stable repository
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu  $(lsb_release -cs)  stable"

## Step 5. Update the package list
sudo apt-get update

## Step 6. Install the latest version of Docker engine
sudo apt-get install docker-ce docker-ce-cli containerd.io

## Step 7. Add user to Docker group
sudo usermod -aG docker $USER

id -nG

## Step 8. Check Docker version
docker --version

## Step 9. Verify the Docker Engine
docker run hello-world


# Docker-Compose
## Step 1. Download the latest version of the Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.25.5/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

## Step 2. Apply executable permissions to the binary
sudo chmod +x /usr/local/bin/docker-compose

## Step 3. Check Docker Compose version
docker-compose --version


# Go
## Step 1. Download the tar file
curl -O https://storage.googleapis.com/golang/go1.12.9.linux-amd64.tar.gz

## Step 2. Extract the tar file
tar -xvf go1.12.9.linux-amd64.tar.gz

## Step 3. Move the go directory
sudo mv go /usr/local

## Step 4. Update environment variables
nano ~/.profile

export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

source ~/.profile

## Step 5. Check go version
go version


# Node.js and NPM
## Step 1. Install the Node.js v10.x repository
curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash -

## Step 2. Install nodejs
sudo apt-get install -y nodejs

## Step 3. Check node version
node -v

## Step 4. Check NPM version
npm -v


# Hyperledger Samples, Binaries and Docker Images
## Step 1. Change directory
cd ..

## Step 2. Download Fabric v1.4.7
curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.7 1.4.7 0.4.20

## Step 3. Update environment variable
nano ~/.profile

export PATH=/home/ubuntu/hyperledger/fabric-samples/bin:$PATH

source ~/.profile
