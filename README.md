# Wire Web Application Demo

**What is this?** This is a simple webserver that just serves a visitor counter and its uptime seconds.
It is built using Go, Postgres, and structured using dependency injection via 
[Google Wire](https://github.com/google/wire). 

**Buy why though?** Because it serves as a good example for the future, for myself and others. The
DI "modules" (config, log, db, etc) are fully-fledged and feature good unit testing and environment-based 
setup. They are ready to be copy-pasted to other projects, or used as reference points.

Some rationale and discussion behind how/why this was put together can be found [in a blog post here](https://medium.com/@fsufitch/dependency-injection-and-testability-in-a-go-webservice-a91d0e5469dd).

## Setup and running

**Building.** The code is built using Go 1.12 and Go modules. All you need to do to build it is:

    go build

To re-generate `wire_gen.go`, use:

    go generate

**Testing.** To run the unit tests and see coverage:

    go test ./... -cover

On Windows, ignore any UAC port opening warning that comes up.

**Environment.** The runtime is parameterized via environment variables:

- `PORT` - the port to serve the webserver on (optional; default=8080)
- `DATABASE` - the PostgreSQL connection string for its database (required)
- `DEBUG` - true/false toggle for debug logging (optional; default=false)

**Database setup.** The code bootstraps its own schema, so you just need to provide it with a database.
An easy Docker-based way to do this would be:

    docker run -p 5432:5432 postgres:latest
    docker run --net=host -it postgres:latest psql postgres://postgres@localhost:5432
    > CREATE DATABASE demo;

This would result in this connection string to be set as `DATABASE` (note the `sslmode` parameter):

    postgres://postgres@localhost:5432/demo?sslmode=disable

**Endpoints.**

- `/uptime` - displays the seconds since the server got started
- `/counter` - increments and displays a visitor counter from the database
