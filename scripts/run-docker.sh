#! /bin/sh

docker build -t switchcraft .

# Run with .env
docker run --rm -p 8080:8080 --net switchcraft_network --env-file .env switchcraft migrate up

# docker run --rm -p 8080:8080 --net switchcraft_network --env-file .env switchcraft serve