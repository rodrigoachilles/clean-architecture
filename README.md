# clean-architecture

Translations:

* [French](README_fr.md)
* [Portuguese (Brazil)](README_pt_br.md)

## Overview

The project consists in a simple _creation_ and _listing_ of all payment orders. The project was built as a challenge for the Go Expert postgraduate course and was obviously written in the Go language.

A payment order contains the following information:
* **Id** - Order Id, generated automatically by the system.
* **ProductName** - Product name.
* **Price** - Order price.
* **Tax** - Fee to be applied to the order price.
* **FinalPrice** - Final price taking into account the order price and the tax.

## Steps to be executed

### docker compose
First, there is a docker file (docker-compose.yaml) to be executed before starting the system. It will initialize the MySql database and RabbitMQ. The following command can be run from the root of the project:
```bash
docker-compose up -d
```

### migrate up
Then use _migration_ to create the _Order_ table in the MySql database:

```bash
make migrate/up
```

### go run
Finally, in the root of the project, run the **main.go** file, located in the **./cmd/ordersystem** directory, with the following command:

```bash
go run .\cmd\ordersystem\main.go .\cmd\ordersystem\wire_gen.go
```

### clients
To execute the commands on the client side, simply use the two _.http_ files, located in the **./api** directory. These will help you execute the commands directly in the Web, gRPC and GraphQL services:
* create_order.http
* list_orders.http

## Services

The project has 4 services, divided into:

### Web Service (REST)

The web service is configured to respond on port **8000** on localhost.
```bash
http://localhost:8000/
```

### gRPC

The gRPC service is configured to respond on port **50051** on localhost.
```bash
http://localhost:50051/
```

### GraphQL

The GraphQL service is configured to respond on port **8080** on localhost.
```bash
http://localhost:8080/
```

### RabbitMQ

The RabbitMQ service is configured to respond on localhost port **5672** and the administration panel can be accessed on localhost port **15672**.
```bash
http://localhost:15672/
```

## Makefile

* migrate/up - Use migration to create the _Order_ table in the MySql database.
* migrate/down - Uses migration to delete the _Order_ table in the MySql database.
* graphql - Command to run the generation of the GraqhQL schema.
* grpc - Command to run the generation of the gRPC file from the protofile.
* wire - Command to generate the Wire file (dependency injection).
