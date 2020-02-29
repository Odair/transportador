package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"transportador/transportador"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	DbHost     = "db"
	DbUser     = "postgres-dev"
	DbPassword = "mysecretpassword"
	DbName     = "dev"
	Migration  = `CREATE TABLE IF NOT EXISTS Entrega (
		IdEntrega serial PRIMARY KEY,
		IdPedido int NOT NULL,
		DataParaBusca timestamp,
		PrevisaoParaEntrega timestamp,
		EnderecoOrigem text,
		EnderecoDestino text,
		CreatedAt timestamp with time zone DEFAULT current_timestamp,
		UpdatedAt timestamp)`
)

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "transportador",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		conexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
		db, err = sql.Open("postgres", conexao)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

		defer db.Close()

		_, err = db.Query(Migration)
		if err != nil {
			level.Info(logger).Log("msg", "migration failed "+err.Error())
			return
		}

	}

	flag.Parse()
	ctx := context.Background()
	var srv transportador.Service
	{
		repository := transportador.NewRepo(db, logger)

		srv = transportador.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := transportador.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := transportador.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
