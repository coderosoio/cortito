#!/usr/bin/env bash

readonly app=cortito

readonly account_image=cortito_account:latest
readonly shortener_image=cortito_shortener:latest
readonly api_image=cortito_api:latest

readonly nginx_image=cortito_nginx:latest
readonly nginx_dockerfile=./docker/nginx/Dockerfile

readonly bzip_file=${app}-latest.tar.bz2

readonly app_key="~/.ssh/id_rsa"
readonly app_user="core@coderoso.io"

readonly images=('account' 'shortener' 'api')

clear

echo "Building app and images..."
docker rmi -f ${nginx_image} 2>/dev/null
docker build --no-cache -t ${nginx_image} -f ./docker/nginx/Dockerfile .
for image in "${images[@]}"; do
  docker rmi -f ${image}:latest 2>/dev/null
  docker build --no-cache -t ${image}:latest -f ./${image}/docker/${image}/Dockerfile .
done

echo "Saving images..."
docker save ${images[@]} | bzip2 > ${bzip_file}

echo "Uploading docker images..."
scp -i ${app_key} ${bzip_file} docker-compose.yml .env.production ${app_user}:/home/core/apps/cortito
for image in ${images[@]}; do
  scp -i ${app_key} ${image}/config.yml ${image}/config.production.yml ${app_user}:/home/core/apps/cortito/${image}/
done

ssh -T -i ${app_key} ${app_user} << ENDSSH
cd /home/core/apps/cortito
cp .env.production .env
docker rmi ${nginx_image} 2>/dev/null
for image in ${images[@]}; do
  docker rmi ${image} 2>/dev/null
done
bunzip2 --stdout ${bzip_file} | docker load
rm ${bzip_file}
docker-compose down && docker-compose up -d --force-recreate
ENDSSH
