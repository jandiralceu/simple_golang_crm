# CRM

A simple CRM API.

## Run app locally

```
go run cmd/server/main.go
```

## Build

In the root project terminal run:

```
go build -o build/app cmd/server/main.go
```

The build version will be: `build/app`

## Coverage

To see the coverage locally, open the c.out file. This file is generated after the push to remote branch.

```
go tool cover -html="c.out"
```
