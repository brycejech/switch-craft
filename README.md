# SwitchCraft

Multi-tenant feature flag management engine written in Go

# Contributing

## Dev system requirements

- [Go ^v1.23](https://go.dev/)
- [Golang Migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
- [Docker](https://docker.com)

## Initial setup

Copy the contents of the `.env.sample` file into a new `.env` file at the root of the project. Then,
replace the values for each key with the real values that you would like to use at runtime.

### Local Docker dependencies

If you are configuring your environment for local development/testing with the PostgreSQL and
pgAdmin docker containers (`docker-compose.yaml`) for the first time, the values for the `DB_USER`,
`DB_PASS`, `PG_ADMIN_EMAIL`, and `PG_ADMIN_PASS` variables in your `.env` file will be used to
configure and initialize these services.

This provides a very convenient initial setup but you must remember that if you wish to change any
of these values later that you cannot simply update them in your `.env` file. You must first change
them via the proper channels for each container, then you can change them in your `.env` file.
Failing to update these values in this way could result in loss of access to the services in these
containers.

Before running the `docker-compose.yaml` file for the first time you will need to create a couple
volumes and a network.

```sh
# Create Postgres volume
docker volume create switchcraft_pg_data

# Create pgAdmin volume
docker volume create switchcraft_pgadmin_data

# Create docker network
docker network create switchcraft_network
```

After configuring the `.env` file and creating the docker volumes and network you can now run
compose to create and run the local services needed for development/testing.

```sh
docker compose up

# Or run in detached mode
docker compose up -d
```

### Genesis record

An initial genesis record for the first user must be manually created either by using the `psql`
command line tool or a database client such as the pgAdmin instance running in the Docker compose
network.

The genesis user needs a hashed password (see auth module in CLI) set in the `password` field as
well as a value in the `created_by` field. It is easiest to create the user directly in the database
with the hashed password in an initial query, then run a subsequent query that sets the `created_by`
field by using the `id` of the same row. This will make it appear as though the genesis user created
itself.

## Running from Docker container

Services running on the `switchcraft_network` are available on `localhost` at the ports specified in
the `docker-compose.yaml` file.

**IMPORTANT:** If you run the application built from the local `Dockerfile` using `docker run`, then
you must attach the container to the `switchcraft_network` using the `--net` argument. Additionally,
when running the container on a local Docker network you must update your `DB_HOST` environment
variable to reference the PostgreSQL container name (found in the `docker-compose.yaml` file),
rather than `localhost`. Examples can be found in the `scripts/run-docker.sh` file.

## Running locally

The application can be run locally using the `go run` command or pre-compiled with `go build` and
then run by calling the executable produced.

Most of the application functionality such as database migrations, running of servers, and core
functions are exposed via a CLI.

```sh
# Build the application for current OS/CPU Arch
go build . -o switchcraft

# Migrate database up
./switchcraft migrate up

# List accounts
./switchcraft account getMany

# Get account by ID
./switchcraft account getOne --id 1

# Get account by UUID
./switchcraft account getOne --uuid d724ed11-5a8b-4cb9-ac3d-fc72b717ba52
```

A more complete `go build` command can be found in `scripts/build.sh`.

Use the `go run` command if you wish to build and run from source rather than pre-compiling, e.g.
`go run . account getOne --id 1`
