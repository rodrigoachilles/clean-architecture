# clean-architecture

Traduções:

* [Inglês](README.md)
* [Francês](README_fr.md)

## Visão geral

O projeto consiste em uma simples _criação_ e _listagem_ de todas as ordens de pagamentos. O projeto foi construído como desafio do curso da pós-graduação em Go Expert sendo escrito, obviamente, na linguagem Go. 

Uma ordem de pagamento contém a seguinte informação: 
* **Id** - Id da ordem gerado automaticamente pelo sistema.
* **ProductName** - Nome do produto. 
* **Price** - Preço da ordem.
* **Tax** - Taxa a ser aplicada ao preço da ordem.
* **FinalPrice** - Preço final considerando o preço da ordem e a taxa.

## Passos a serem executados

### docker compose
Existe um arquivo de docker (docker-compose.yaml) a ser executado antes de iniciar o sistema. Ele inicializará o bando de dados MySql e o RabbitMQ. Na raiz do projeto pode ser executado o seguinte comando:
```bash
docker-compose up -d
```

### migrate up
Após isso, utilise a _migration_ de criação da tabela _Order_ no banco de dados MySql:

```bash
make migrate/up
```

### go run
E por fim, na raiz do projeto, execute o arquivo **main.go**, localizados no diretório **./cmd/ordersystem**, com o seguinte comando:

```bash
go run .\cmd\ordersystem\main.go .\cmd\ordersystem\wire_gen.go
```

### clientes
Para executar os comandos do lado do cliente, basta utilizar os dois arquivos _.http_, localizados no diretório **./api**. Eles auxiliarão para executar os comandos diretamente nos serviços Web, gRPC e GraphQL:
* create_order.http
* list_orders.http

## Serviços

O projeto tem 4 serviços, divididos em:

### Serviço Web (REST)

O serviço Web está configurado para responder na porta **8000** do localhost.
```bash
http://localhost:8000/
```

### gRPC

O serviço do gRPC esta configurado para responder na porta **50051** do localhost.
```bash
http://localhost:50051/
```

### GraphQL

O serviço do GraphQL esta configurado para responder na porta **8080** do localhost.
```bash
http://localhost:8080/
```

### RabbitMQ

O serviço de mensageria RabbitMQ esta configurado para responder na porta **5672** do localhost e o painel de administração pode ser acessado na porta **15672** do localhost.
```bash
http://localhost:15672/
```

## Makefile

* migrate/up - Utiliza a migration para criar a tabela _Order_ no banco de dados MySql.
* migrate/down - Utiliza a migration para deletar a tabela _Order_ no banco de dados MySql.
* graphql - Comando para executar a geração do schema do GraqhQL. 
* grpc - Comando para executar a geração do arquivo gRPC a partir do protofile.
* wire - Comando para executar a geração do arquivo Wire (Injeção de dependências).
