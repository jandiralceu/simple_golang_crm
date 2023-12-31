![CI-CD Status](https://github.com/jandiralceu/simple_golang_crm/actions/workflows/main.yml/badge.svg?event=push) [![Coverage Status](https://coveralls.io/repos/github/jandiralceu/simple_golang_crm/badge.svg)](https://coveralls.io/github/jandiralceu/simple_golang_crm)

# CRM Rest-API

This simple CRM Rest-API is the final project from Udacity [Golang course](https://www.udacity.com/course/golang--cd11970).

All commands bellow, must be typed on project root directory.

## Install

```
go mod download
```

## Run

```shell
go run cmd/server/main.go
```

## Build

```shell
go build cmd/crm/main.go
```

The build version will be located at root folder as: `main`

## Test and Coverage

The command bellow will test and generate the coverage file.

```shell
go test -coverprofile=c.out ./...
```

To see the coverage type the command bellow:

```shell
go tool cover -html="c.out"
```

Evertime the project is submitted to github, the pipeline generate the coverage and upload it to coveralls. YOu can see the results on the badge bellow.

[![Coverage Status](https://coveralls.io/repos/github/jandiralceu/simple_golang_crm/badge.svg)](https://coveralls.io/github/jandiralceu/simple_golang_crm)
