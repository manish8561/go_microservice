# ![Microservice]

Microservice with Express Gateway, Typescript, Docker and Go Lang Services

> ### Golang/Gin codebase containing real world examples (CRUD, auth, advanced patterns, etc) that adheres to the [Autocompound] spec and API.


This codebase was created to demonstrate a fully fledged fullstack application built with **Golang/Gin** including Mongodb, CRUD operations, authentication, routing, pagination, and more.


# Directory structure
```
docker_backend
|
gateway_service // express gateway service
|
event_service //event service in typescript to fetch events
|
contract_service // contract service in typscript to calculate APR
|   
user // user service
├── server.go
├── common
│   ├── utils.go        //small tools function
│   └── database.go     //DB connect manager
├── users
|   ├── models.go       //data models define & DB operation
|   ├── routers.go      //business logic & router binding
|   ├── middlewares.go  //put the before & after logic of handle request
|   └── validators.go   //form/json checker
├── ...
farm // farm service
├── server.go
├── common
│   ├── utils.go        //small tools function
│   └── database.go     //DB connect manager
├── farms
|   ├── models.go       //data models define & DB operation
|   ├── routers.go      //business logic & router binding
|   ├── middlewares.go  //put the before & after logic of handle request
|   └── validators.go   //form/json checker
├── pricefeeds  // get price from coingeeko
├── stakes      // CRUD for staking contracts
├── contracts // for contract call  
├── helloworld // for grpc to get user data  
├── ...
...
```

# Getting started

## Install Golang

Make sure you have Go 1.13 or higher installed.

https://golang.org/doc/install

## Environment Config

Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).


## Install Dependencies
From the project root, run:
```
go build ./...
go test ./...
go mod tidy
```

## Testing
From the project root, run:
```
go test ./...
```
or
```
go test ./... -cover
```
or
```
go test -v ./... -cover
```
depending on whether you want to see test coverage and how verbose the output you want.

## Todo
- More elegance config
- Test coverage (common & users 100%, article 0%)
- ProtoBuf support
- Code structure optimize (I think some place can use interface)