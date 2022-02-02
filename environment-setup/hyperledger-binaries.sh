# Hyperledger Samples, Binaries and Docker Images
## Step 1. Change directory
cd ..

## Step 2. Download Fabric v2.2.0 (LTS)
curl -sSL https://bit.ly/2ysbOFE | bash -s -- 2.2.0 1.4.9

## Step 3. Update environment variable
nano ~/.profile

echo 'export PATH=/home/ubuntu/hyperledger/fabric-samples/bin:$PATH' >>~/.profile

source ~/.profile