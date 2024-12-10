# docker-training-exam

This repository contains applications to Dockerize with instructions on how they work. It's a very simple todolist
application that let you create and fetch todos.

# Context

## frontend

The frontend part of this application has been written using VueJS

## backend

The backend part of this application has been written using Golang.

It starts a simple HTTP server that listen on a port defined in a configuration file read by the application.

You have an example of that configuration file [here](backend/configs/app.local.yml).

3 routes are exposed at the moment:
1. `GET /health`: returns "ok"
2. `GET /api/v1/todos`: return all created todos
3. `POST /api/v1/todos`: create a new todo. JSON is expected as input.

```json
{
  "author": "Frodo Baggins",
  "text": "Throw the unique in Amon Amarth"
}
```

`todos` are stored in a `PostgreSQL` database for persistence.

Along with that configuration file, the application use environment variables to configure itself:

| Variable name           | Description                                                                                                                                          |
|-------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------|
| APP_ENVIRONMENT         | Defines if the logger should be configured for production or local development. If the value is "local" logs will be formatted to be human readable. |
| APP_CONFIGURATION_FILE  | Filepath to the application configuration file. If not set, the application will look for it's config in `$(pwd)/configs/app.yml`                    |

## database

The backend is using `PostgreSQL` has it's storage system to persist data.

# Exercise

Extend the `compose.yaml` file to deploy both `frontend`, `backend` and `postgres` services. Write required `Dockerfile`'s
to build the backend and the frontend image.

The compose file already contains a `service` that runs database migration to create SQL objects (tables, ...)

Here's a small ordered list to help you in your adventure:
- [ ] write the `db` service in the `compose.yaml` file to start a PostgresSQL container. Details are [here](https://hub.docker.com/_/postgres)
- [ ] write the **backend** `Dockerfile` to containerize it.
  - [ ] after building the image, try to run it locally to see if it works
  - [ ] make sure the image size is small, be efficient!
  - [ ] extra points if you write the Dockerfile in such a way that we profit smartly from it's caching system
- [ ] add the `backend` service to the `compose.yaml` file
  - [ ] tell `compose` how to build your service
  - [ ] expose the service to port `8080` of your machine
  - [ ] configure the logger in "local" mode
  - [ ] the application should read its configuration file from `/app/config.yml`
  - [ ] mount its configuration file to `/app/config.yml`
  - [ ] we want to be able to fetch the application logs from your machine

Small reminders/hints for building the golang app:
- use the image `golang:1.23` to have access to go building tools
- to build a go application, you can run `CGO_ENABLED=0 go build -o backend *.go`
- **(optional - extra points only)** if you want to install dependencies only, you will need `go.mod` and `go.sum` files and run `go mod download -x`

Small reminder about Docker compose:
- environment variables such as `${POSTGRES_USER}` are read from a file called `.env`...

<details>
  <summary>Good luck</summary>
  <img alt="good luck gif" src="https://i.giphy.com/media/v1.Y2lkPTc5MGI3NjExZnNqZDV1Y2M4ZDh4a3l0dmozeDlsd2xzdmJ1d3VyMHF0dGRvY2JvZSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/3oeSAz6FqXCKuNFX6o/giphy.gif"/>
</details>
