#!/usr/bin/env bash

readonly app=cortito

readonly bzip_file=${app}-latest.tar.bz2

readonly app_key="~/.ssh/id_rsa"
readonly app_user="core@coderoso.io"

clear

echo "Building app and images..."
# docker rmi -f coderoso/cortito_nginx 2>/dev/null
# docker rmi -f coderoso/cortito_account:latest 2>/dev/null
# docker rmi -f coderoso/cortito_shortener:latest 2>/dev/null
# docker rmi -f coderoso/cortito_api:latest 2>/dev/null
# docker rmi -f coderoso/cortito_frontend:latest 2>/dev/null
# docker rmi -f coderoso/cortito_web:latest 2>/dev/null

docker build -t coderoso/cortito_nginx -f ./docker/nginx/Dockerfile .
docker build -t coderoso/cortito_account -f ./account/docker/account/Dockerfile .
docker build -t coderoso/cortito_shortener -f ./shortener/docker/shortener/Dockerfile .
docker build -t coderoso/cortito_api -f ./shortener/docker/shortener/Dockerfile .
docker build -t coderoso/cortito_frontend -f ./frontend/docker/frontend/Dockerfile .
docker build -t coderoso/cortito_web -f ./web/docker/web/Dockerfile .

echo "Saving images..."
docker save coderoso/cortito_nginx \
            coderoso/cortito_account \
            coderoso/cortito_shortener \
            coderoso/cortito_frontend \
            coderoso/cortito_web \
            coderoso/cortito_api | bzip2 > ${bzip_file}

echo "Uploading docker images..."
scp -i ${app_key} ${bzip_file} docker-compose.production.yml .env.production ${app_user}:/home/core/apps/cortito
scp -i ${app_key} ./account/config.yml ./account/config.production.yml ${app_user}:/home/core/apps/cortito/account/
scp -i ${app_key} ./shortener/config.yml ./shortener/config.production.yml ${app_user}:/home/core/apps/cortito/shortener/
scp -i ${app_key} ./api/config.yml ./api/config.production.yml ${app_user}:/home/core/apps/cortito/api/
scp -i ${app_key} ./web/config.yml ./web/config.production.yml ${app_user}:/home/core/apps/cortito/web/

ssh -T -i ${app_key} ${app_user} << ENDSSH
cd /home/core/apps/cortito

mv docker-compose.production.yml docker-compose.yml
mv .env.production .env

docker rmi -f coderoso/cortito_nginx:latest 2>/dev/null
docker rmi -f coderoso/cortito_account:latest 2>/dev/null
docker rmi -f coderoso/cortito_shortener:latest 2>/dev/null
docker rmi -f coderoso/cortito_api:latest 2>/dev/null
docker rmi -f coderoso/cortito_web:latest 2>/dev/null
docker rmi -f coderoso/cortito_frontend:latest 2>/dev/null

bunzip2 --stdout ${bzip_file} | docker load

rm ${bzip_file}

docker-compose down && docker-compose up -d --force-recreate
ENDSSH
