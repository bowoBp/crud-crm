# README #

## Setup

1. Install Go version 1.19
2. Install Mockery version v2.13 or later
3. Use GoLand (recommended)
4. Download dependencies with command `go mod download`
4. Create `.env` file based on `.env.example`

## Run

Use this command to run API app from root directory:

```shell
go run app/api/api.go
```

Use this command to run consumer app from root directory:

```shell
go run app/consumer/consumer.go
```

Use this command to run scheduler app from root directory:

```shell
go run app/scheduler/scheduler.go
```

## Unit Tests

### Generate Mocks

To generate mock, run:

```shell
mockery --all --keeptree --case underscore --with-expecter
```

### Run Unit Tests

To run unit tests:
```shell
go test ./...
```

---

## TODO README

This README would normally document whatever steps are necessary to get your application up and running.

### What is this repository for? ###

* Quick summary
* Version
* [Learn Markdown](https://bitbucket.org/tutorials/markdowndemo)

### How do I get set up? ###

* Summary of set up
* Configuration
* Dependencies
* Database configuration
* How to run tests
* Deployment instructions

### Contribution guidelines ###

* Writing tests
* Code review
* Other guidelines

### Who do I talk to? ###

* Repo owner or admin
* Other community or team contact

# Doumentation API With Swager
# go swag
An example repo for auto generating RESTful API doc using https://github.com/swaggo/swag

### Starting development server
- Copy `.env.example` to `.env` in the same directory
- Update environment variables with credentials for mysql
```
go run main.go
```

### Swagger 2.0 API documents
Run the app, and browse to http://localhost:[SERVER_PORT]/swagger/index.html. Find the Swagger 2.0 Api documents