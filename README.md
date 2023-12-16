# Установка программ на VPS

 1. Vim, Curl, Gpg, Git
``` shell
apt update
apt install -y vim curl gpg git
```

2. Docker Engine
``` shell
apt install -y apt-transport-https ca-certificates curl software-properties-common
curl -fsSL https://download.docker.com/linux/debian/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/debian $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
apt update
apt install -y docker-ce docker-ce-cli containerd.io
systemctl start docker
systemctl enable docker
```


# Работа с Git репозиторием на VPS

1. Авторизация в GitHub
Создать SSH ключ, следуя [документации](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent)

Добавить созданный ключ на GitHub, следуя [документации](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

2. Клонирование проекта с удаленного репозитория GitHub
``` shell
git clone git@github.com:spotich/routes.git
```


# Запуск приложения

``` shell
cd routes
docker build -t routes .

```
