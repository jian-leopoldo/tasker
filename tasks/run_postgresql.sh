#!/bin/bash
docker run -d \
    --name postgres-prisma \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=mysecretpassword \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v /custom/mount:/var/lib/postgresql/data \
    postgres
