#!/bin/bash

docker pull postgres:15.1
if [ ! "$(docker ps -q -f name=pgsql1)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=pgsql1)" ]; then
        docker rm pgsql1
    fi
    docker run --name=pgsql1 -p 5432:5432 -v "/Users/rmjv/Documents/go/src/vivaopt/mntdata":/var/lib/postgresql/data -e POSTGRES_PASSWORD=password -e POSTGRES_DB=shortlink -d postgres:15.1
    # ss -tulpn | grep 5432
fi
