For later:

    docker run -p 5432:5432 postgres:latest
    docker run --net=host -it postgres:latest psql postgres://postgres@localhost:5432
    > CREATE DATABASE demo;
    DATABASE=postgres://postgres@localhost:5432/demo?sslmode=disable ./testable-web-demo.exe