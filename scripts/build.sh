#! /bin/sh

CGO_ENABLED=0 go build -a -ldflags="-w -s" .