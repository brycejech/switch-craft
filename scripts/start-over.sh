#! /bin/sh

go run . migrate down;
go run . migrate up;
go run . seed all --dataFile=seed.json