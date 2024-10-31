#! /bin/sh

docker build -t switchcraft .

# Migrate database up, use local .env file
docker run --rm -p 8080:8080 --net switchcraft_network --env-file .env switchcraft migrate up

# List accounts
# docker run --rm -p 8080:8080 --net switchcraft_network --env-file .env switchcraft account getMany