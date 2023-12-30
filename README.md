![Github Actions Status](https://github.com/jandiralceu/simple_golang_crm/actions/workflows/main.yml/badge.svg)

# CRM

This simple CRM API is the final project from Udacity Golang course.

## Run app locally

```
go run cmd/server/main.go
```

## Build

In the root project terminal run:

```
go build -o build/app cmd/server/main.go
```

The build version will be located at: `build/app`

## Coverage

To see the coverage locally, open the c.out file. This file is generated after the push to remote branch.

```
go tool cover -html="c.out"
```
