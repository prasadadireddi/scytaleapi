# Go RESTful API for Workload Registration

This framework implement a registration entry management SPA using Go

It encourages writing clean and idiomatic Go code. 

The provided registration API with ability to::

* RESTful endpoints in the widely accepted format
* Standard CRUD operations of a database table
* Create Entries
* Delete Entries
* Edit Entries
* Display Entries
* Data validation
* Full test coverage

The kit uses the following Go packages which can be easily replaced with your own favorite ones
since their usages are mostly localized and abstracted. 

* Routing: [ozzo-routing](https://github.com/go-ozzo/ozzo-routing)
* Database access: [ozzo-dbx](https://github.com/go-ozzo/ozzo-dbx)
* Database migration: [golang-migrate](https://github.com/golang-migrate/migrate)
* Data validation: [ozzo-validation](https://github.com/go-ozzo/ozzo-validation)
* Logging: [zap](https://github.com/uber-go/zap)
* JWT: [jwt-go](https://github.com/dgrijalva/jwt-go)

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The kit requires **Go 1.13 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the kit without setting up your
own database server. The kit requires **Docker 17.05 or higher** for the multi-stage build support.

After installing Go and Docker, run the following commands to start experiencing this starter kit:

```shell
# download the starter kit
git clone https://github.com/qiangxue/go-rest-api.git

cd go-rest-api

# start a PostgreSQL database server in a Docker container
make db-start

# seed the database with some test data
make testdata

# run the RESTful API server
make run

# or run the API server with live reloading, which is useful during development
# requires fswatch (https://github.com/emcrisostomo/fswatch)
make run-live
```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

* `GET /healthcheck`: a healthcheck service provided for health checking purpose (needed when implementing a server cluster)
* `POST /v1/login`: authenticates a user and generates a JWT
* `GET /v1/albums`: returns a paginated list of the albums
* `GET /v1/albums/:id`: returns the detailed information of an album
* `POST /v1/albums`: creates a new album
* `PUT /v1/albums/:id`: updates an existing album
* `DELETE /v1/albums/:id`: deletes an album

Try the URL `http://localhost:8080/healthcheck` in a browser, and you should see something like `"OK v1.0.0"` displayed.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the following 
more complex scenarios:

```shell
# authenticate the user via: POST /v1/login
curl -X POST -H "Content-Type: application/json" -d '{"username": "demo", "password": "pass"}' http://localhost:8080/v1/login
# should return a JWT token like: {"token":"...JWT token here..."}

# with the above JWT token, access the album resources, such as: GET /v1/albums
curl -X GET -H "Authorization: Bearer ...JWT token here..." http://localhost:8080/v1/albums
# should return a list of album records in the JSON format
```

To use the starter kit as a starting point of a real project whose package name is `github.com/abc/xyz`, do a global 
replacement of the string `github.com/qiangxue/go-rest-api` in all of project files with the string `github.com/abc/xyz`.


## Project Layout

The starter kit uses the following project layout:
 
```
.
├── Dockerfile
├── Dockerfile.server
├── README.md
├── api
│   ├── controllers
│   │   └── basic_controller.go
│   ├── database
│   │   └── db.go
│   ├── models
│   │   └── workload.go
│   ├── repository
│   │   ├── crud
│   │   │   └── repository_workload_crud.go
│   │   └── workload.go
│   ├── responses
│   │   └── json.go
│   ├── router
│   │   ├── router.go
│   │   └── routes
│   │       ├── basic.go
│   │       └── routes.go
│   ├── server.go
│   └── utils
│       ├── channels
│       │   └── channels.go
│       └── console
│           └── console.go
├── auto
│   ├── data.go
│   └── load.go
├── config
│   └── config.go
├── docker-compose.yaml
├── docker-compose.yaml.orig
├── go.mod
├── go.sum
├── main.go
└── main_test.go

```

The top level directories `cmd`, `internal`, `pkg` are commonly found in other popular Go projects, as explained in
[Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Within `internal` and `pkg`, packages are structured by features in order to achieve the so-called
[screaming architecture](https://blog.cleancoder.com/uncle-bob/2011/09/30/Screaming-Architecture.html). For example, 
the `album` directory contains the application logic related with the album feature. 

Within each feature package, code are organized in layers (API, service, repository), following the dependency guidelines
as described in the [clean architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).