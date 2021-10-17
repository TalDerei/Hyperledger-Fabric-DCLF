# Hyperledger Samples, Binaries and Docker Images
## Step 1. Change directory
cd ..

## Step 2. Download Fabric v1.4.7
curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.7 1.4.7 0.4.20

## Step 3. Update environment variable
nano ~/.profile

echo 'export PATH=/home/ubuntu/hyperledger/fabric-samples/bin:$PATH' >>~/.profile

source ~/.profile