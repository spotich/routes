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


# Подготовка VPS

## Git
1. Авторизация в GitHub
Создать SSH ключ, следуя [документации](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent). Добавить созданный ключ на GitHub, следуя [документации](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

2. Клонирование проекта с удаленного репозитория GitHub
``` shell
git clone git@github.com:spotich/routes.git
```

## Docker
3. Установка образов
``` shell
docker pull spotich/routes
docker pull mysql
docker pull phpmyadmin
```

4. Создание сети
``` shell
docker network create routes
```


# Запуск приложения

### Создание контейнеров

1. Приложение
``` shell
cd routes
docker run \
    --name routes-app \
    --network routes \
    -dp 80:80 \
    spotich/routes
```

2. MySQL
``` shell
mkdir -p ./storage/mysql
docker run \
    --name routes-db \
    --network routes \
    -dp 3306:3306 \
    -e MYSQL_ROOT_PASSWORD=1234 \
    -v ./storage/mysql:/var/lib/mysql \
    mysql
```

3. PhpMyAdmin
``` shell
docker run \
    --name routes-pma \
    --network routes \
    -dp 8080:8080 \
    -e PMA_HOST=routes-db \
    phpmyadmin
```
