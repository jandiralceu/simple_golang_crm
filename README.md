![CI-CD Status](https://github.com/jandiralceu/simple_golang_crm/actions/workflows/main.yml/badge.svg?event=push)

# Simple CRM

This simple CRM API is the final project from Udacity Golang course.

## Run app locally

```
go run cmd/server/main.go
```

## Build

In the root project terminal run:

```
go build cmd/crm/main.go
```

The build version will be located at root folder as: `main`

## Coverage

To see the coverage locally, open the c.out file. This file is generated after the push to remote branch.

```
go tool cover -html="c.out"
```
