# Tiket API

Tiket API for online Loket reservation and implementing Clean Architecture for seperation of concern.

## Getting Started

### Prerequisites

Things you need to have and installed

```
MySQL v5.7
Go v1.13 or latest
```

### Installing

Setup database

![alt text](https://github.com/kadekchrisna/tiket/blob/master/assets/loket.jpg?raw=true)

```
create scheme with name loket

import script tiket.sql
```

Copy Environtment Variables

```
cp local.env .env
```

Installing dependencies

```
go to project directory, run

go get -u ./...
```

## Running the server

How to run the server

```
./run.sh
```
