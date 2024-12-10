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

| Variable name           | Description                                                                                                                                        |
|-------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| APP_ENVIRONMENT         | Defines if the logger should be configured for production or local development. If the value is "dev" logs will be formatted to be human readable. |
| APP_CONFIGURATION_FILE  | Filepath to the application configuration file. If not set, the application will look for it's config in `$(pwd)/configs/app.yml`                  |

## database

The backend is using `PostgreSQL` has it's storage system to persist data.

# Exercise

[ ] todo