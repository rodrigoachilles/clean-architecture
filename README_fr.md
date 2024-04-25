# clean-architecture

Traductions:

* [Anglais](README.md)
* [Portugais (Brésil)](README_pt_br.md)

## Introduction

Le projet consiste en une simple _création_ et _liste_ de tous les ordres de paiement. Le projet a été conçu comme un défi pour le cours de troisième cycle Go Expert et a évidemment été écrit en langage Go.

Un ordre de paiement contient les informations suivantes :
* **Id** - Identité de l'ordre, générée automatiquement par le système.
* **ProductName** - Nom du produit.
* **Price** - Prix de la commande.
* **Tax** - Taxe à appliquer au prix de la commande.
* **FinalPrice** - Prix final tenant compte du prix de la commande et de la taxe.

## Étapes à exécuter

### docker compose
Tout d'abord, un fichier docker (docker-compose.yaml) doit être exécuté avant de démarrer le système. Il initialisera la base de données MySql et RabbitMQ. La commande suivante peut être exécutée depuis la racine du projet :
```bash
docker-compose up -d
```

### migrate up
Utilisez ensuite _migration_ pour créer la table _Order_ dans la base de données MySql :

```bash
make migrate/up
```

### go run
Enfin, à la racine du projet, exécutez le fichier **main.go**, situé dans le répertoire **./cmd/ordersystem**, avec la commande suivante :

```bash
go run .\cmd\ordersystem\main.go .\cmd\ordersystem\wire_gen.go
```

### clients
Pour exécuter les commandes du côté client, il suffit d'utiliser les deux fichiers _.http_ situés dans le répertoire **./api**. Ils vous aideront à exécuter les commandes directement dans les services Web, gRPC et GraphQL :
* create_order.http
* list_orders.http

## Services

Le projet a 4 services, divisés en :

### Web Service (REST)

Le service web est configuré pour répondre sur le port **8000** sur localhost.
```bash
http://localhost:8000/
```

### gRPC

Le service gRPC est configuré pour répondre sur le port **50051** sur localhost.
```bash
http://localhost:50051/
```

### GraphQL

Le service GraphQL est configuré pour répondre sur le port **8080** sur localhost.
```bash
http://localhost:8080/
```

### RabbitMQ

Le service RabbitMQ est configuré pour répondre sur le port **5672** de localhost et le panneau d'administration est accessible sur le port **15672** de localhost.
```bash
http://localhost:15672/
```

## Makefile

* migrate/up - Utilise la migration pour créer la table _Order_ dans la base de données MySql.
* migrate/down - Utilise la migration pour supprimer la table _Order_ dans la base de données MySql.
* graphql - Commande permettant d'exécuter la génération du schéma GraqhQL.
* grpc - Commande permettant de générer le fichier gRPC à partir du protofile.
* wire - Commande permettant de générer le fichier Wire (injection de dépendances).

