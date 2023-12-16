# Установка программ на VPS
## 1. Docker Engine
```
apt update
apt install -y curl
apt install -y apt-transport-https ca-certificates curl software-properties-common
apt install -y gpg
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
apt update
apt install -y docker-ce docker-ce-cli containerd.io
systemctl start docker
systemctl enable docker
```
