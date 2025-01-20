# Add dependencies
sudo apt install make

# Add Docker's official GPG key:

sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:

echo \
 "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
 $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
 sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
sudo docker run hello-world

# post-intall

sudo groupadd docker
sudo usermod -aG docker $USER
newgrp docker
docker run hello-world

# manage working space

mkdir working
cd working
mkdir infras
mkdir repos
cd infras

# clone repose

git clone https://github.com/kingstonduy/debezium.git
git clone https://github.com/kingstonduy/redis.git
git clone https://github.com/kingstonduy/elk-thesis.git

cd ..
cd repos
git clone https://github.com/kingstonduy/thesis-backend.git

# install golang

curl -OL https://golang.org/dl/go1.22.10.linux-amd64.tar.gz
sudo tar -C /usr/local -xvf go1.22.10.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
source ~/.profile
go version
