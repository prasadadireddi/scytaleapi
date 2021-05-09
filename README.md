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
* Test coverage

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. The API requires **Go 1.15 or above**.

[Docker](https://www.docker.com/get-started) is also needed if you want to try the API without setting up your
own database server. The app requires **Docker 17.05 or higher** for the multi-stage build support.

[Spiffe/Spire](https://spiffe.io/docs/latest/spire-about/spire-concepts/) is needed for one of the API validation. This app assumes the setup is installed and configured on remote machine before perfmoring `POST /api/v1/svid/validate` request.

After installing Go and Docker, run the following commands to start experiencing this RESTful API for workload registration:

```shell
# download the starter API code
git clone https://github.com/prasadadireddi/scytaleapi.git

cd scytaleapi

# create .env file for Environment details with below contents
$ cat .env 
API_PORT=8080

# DATABASE CONFIG
DB_DRIVER=postgres
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=admin@123
DB_NAME=postgres


# run the RESTful API server with unit tests
docker compose up

# or run the API server with live, which is useful to perform curl requests
docker compose -f docker-compose-server.yaml up

# shutdown server after unit tests are completed
docker compose down

```

At this time, you have a RESTful API server running at `http://127.0.0.1:8080`. It provides the following endpoints:

# SVID Validation
* `POST /api/v1/svid/validate`: service to validate SAN of SPIFFEID configuration parameter passed is matches the registered SVID

# Workload Entry registration
* `GET /api/v1/workloads`: service provided to get the workloads registered
* `GET /api/v1/workloads/sorted`: service provided to get the workloads in sorted order based on SPIFFEID
* `GET /api/v1/workload/{selector}`: service to get workload(s) based on selector
* `POST /api/v1/workload`: service to register new workload entry
* `PUT /api/v1/workload/{spiffeid}`: service to update existing workload entry
* `PUT /api/v1/workload/{spiffeid}/{selector}`: service to update selectors for particular workload
* `DELETE /api/v1/workload/{spiffeid}`: deletes an entry from registered workloads
* `DELETE /api/v1/workload/{spiffeid}/{selector}`: deletes a selector from a particular workload


Try the URL `http://localhost:8080/api/v1/workloads` in a browser, and you should see something like `"OK v1.0.0"` displayed.

If you have `cURL` or some API client tools (e.g. [Postman](https://www.getpostman.com/)), you may try the below scenarios:

```shell
# validate spiffeid parameter passed matches with SAN of the registered workload SVID. return 200 if matches, and 401 if it doesn't
curl --header "Content-Type: application/json" --request POST --data '{"spiffeid":"https://example.com/service"}' http://10.0.1.1:8080/api/v1/svid/validate

Assumption: The spire-agent is running on 10.0.1.1 and is attested with spire-server.

# register new workload entry via: POST /api/v1/workload
curl --header "Content-Type: application/json" --request POST --data '{"spiffeid":"spiffe://trust-domain-name/path","selectors":"app:demo"}' http://127.0.0.1:8080/api/v1/workload

# update existing workload entry via: PUT /api/v1/workload/{spiffeid}
curl --header "Content-Type: application/json" --request PUT --data '{"selectors":["app:demo"]}' http://127.0.0.1:8080/api/v1/workload/test

# delete existing workload entry via: DELETE /api/v1/workload/{spiffeid}
curl --header "Content-Type: application/json" --request DELETE  http://127.0.0.1:8080/api/v1/workload/test

```

## Project Layout

The Workload resgitration API service uses the following project layout:
 
```
.
├── Dockerfile
├── Dockerfile.server
├── README.md
├── api
│   ├── controllers
│   │   ├── basic_controller.go
│   │   └── svid_controller.go
│   ├── database
│   │   └── db.go
│   ├── models
│   │   └── workload.go
│   ├── repository
│   │   ├── crud
│   │   │   ├── repository_svid_crud.go
│   │   │   └── repository_workload_crud.go
│   │   ├── svid.go
│   │   └── workload.go
│   ├── responses
│   │   └── json.go
│   ├── router
│   │   ├── router.go
│   │   └── routes
│   │       ├── basic.go
│   │       ├── routes.go
│   │       └── svid.go
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
├── docker-compose-server.yaml
├── docker-compose.yaml
├── go.mod
├── go.sum
├── main.go
└── main_test.go

```

## Reference

