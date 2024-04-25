package main

import (
	"database/sql"
	"fmt"
	graphqlHandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rodrigoachilles/clean-architecture/configs"
	"github.com/rodrigoachilles/clean-architecture/internal/event/handler"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/graph"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/grpc/pb"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/grpc/service"
	"github.com/rodrigoachilles/clean-architecture/internal/infra/web/webserver"
	"github.com/rodrigoachilles/clean-architecture/internal/usecase"
	"github.com/rodrigoachilles/clean-architecture/pkg/events"
	"github.com/rodrigoachilles/clean-architecture/pkg/log"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Logger.Info().Msg("Starting the application ...")

	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Logger.Panic().Err(err).Msg("Error loading config")
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	db, err := sql.Open(cfg.DBDriver, dataSourceName)
	if err != nil {
		log.Logger.Panic().Err(err).Msg("Error connecting to database")
	}
	defer db.Close()

	rabbitMQChannel := getRabbitMQChannel(cfg)
	eventDispatcher := events.NewEventDispatcher()
	_ = eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	listOrderUseCase := NewListOrdersUseCase(db)

	startWebServer(eventDispatcher, db, cfg)
	startGrpcServer(createOrderUseCase, listOrderUseCase, cfg)
	startGraphQLServer(createOrderUseCase, listOrderUseCase, cfg)
}

func startGraphQLServer(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase, cfg *configs.Conf) {
	srv := graphqlHandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		ListOrdersUseCase:  *listOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Logger.Info().Msg(fmt.Sprintf("Starting GraphQL server on port %s ...", cfg.GraphQLServerPort))
	err := http.ListenAndServe(":"+cfg.GraphQLServerPort, nil)
	if err != nil {
		log.Logger.Panic().Err(err).Msg("Error connecting to Web Service")
	}
}

func startGrpcServer(createOrderUseCase *usecase.CreateOrderUseCase, listOrdersUseCase *usecase.ListOrdersUseCase, cfg *configs.Conf) {
	grpcServer := grpc.NewServer()
	orderService := service.NewOrderService(*createOrderUseCase, *listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, orderService)
	reflection.Register(grpcServer)

	log.Logger.Info().Msg(fmt.Sprintf("Starting gRPC server on port %s ...", cfg.GRPCServerPort))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServerPort))
	if err != nil {
		log.Logger.Panic().Err(err).Msg("Error connecting to gRPC Server")
	}
	go grpcServer.Serve(lis)
}

func startWebServer(eventDispatcher *events.EventDispatcher, db *sql.DB, cfg *configs.Conf) {
	ws := webserver.NewWebServer(cfg.WebServerPort)
	orderHandler := NewOrderHandler(db, eventDispatcher)
	ws.AddHandler("/order", orderHandler.Handle)
	log.Logger.Info().Msg(fmt.Sprintf("Starting Web Service on port %s ...", cfg.WebServerPort[1:]))
	go ws.Start()
}

func getRabbitMQChannel(cfg *configs.Conf) *amqp.Channel {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.MQUser, cfg.MQPassword, cfg.MQHost, cfg.MQPort)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Logger.Panic().Err(err).Msg("Error connecting to RabbitMQ, check the environment variables")
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Panic().Err(err).Msg("Error connecting to RabbitMQ Channel")
	}
	return ch
}
