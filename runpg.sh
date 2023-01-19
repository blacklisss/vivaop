#!/bin/bash

docker pull postgres:15.1
if [ ! "$(docker ps -q -f name=pgsql1)" ]; then
  if [ "$(docker ps -aq -f status=exited -f name=pgsql1)" ]; then
    docker rm pgsql1
  fi
  docker run --name=pgsql1 \
  -P -p 127.0.0.1:5432:5432 \
  -v "$(pwd)/mntdata":/var/lib/postgresql/data \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=vivaop \
  -d -v "$(pwd)/internal/infrastructure/db/pgstore/migration":/docker-entrypoint-initdb.d \
  postgres:15.1
  # --network host \
  # ss -tulpn | grep 5432
fi
