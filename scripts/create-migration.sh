#! /bin/sh

if [ -z "$1" ]
  then
    echo "Must pass migration name"
    exit 1
fi

migrate create -ext sql -dir repository/queries/migrations -seq $1