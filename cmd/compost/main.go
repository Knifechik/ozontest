package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sipki-tech/database/connectors"
	"gopkg.in/yaml.v3"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"ozon_test_compost/cmd/compost/internal/adapters/in_memory"
	"ozon_test_compost/cmd/compost/internal/adapters/repo"
	"ozon_test_compost/cmd/compost/internal/api/graph"
	"ozon_test_compost/cmd/compost/internal/app"
	"syscall"
)

const defaultPort = "8080"
const migrateDir = "migrate"

type (
	config struct {
		DB dbConfig `yaml:"db"`
	}

	dbConfig struct {
		MigrateDir string         `yaml:"migrate_dir"`
		Driver     string         `yaml:"driver"`
		Postgres   repo.Connector `yaml:"postgres"`
	}
)

var (
	postgresDBConf = connectors.PostgresDB{
		User:     "user123",
		Password: "pass123",
		Host:     "postgres",
		Port:     5432,
		Database: "postgres",
		Parameters: &connectors.PostgresDBParameters{
			Mode: connectors.PostgresSSLDisable},
	}
	flagRepo = flag.String("repotype", "postgres", "what type of repository to use")
)

func main() {
	flag.Parse()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	defer cancel()

	cfg := configRead()

	var r app.Repo

	switch *flagRepo {
	case "postgres":
		myDB, err := repo.New(ctx, repo.Config{
			Postgres:   cfg.DB.Postgres,
			MigrateDir: cfg.DB.MigrateDir,
			Driver:     cfg.DB.Driver,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			err := myDB.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		r = myDB
	case "in_memory":
		r = in_memory.New()

	}

	myApp := app.New(r)

	mux := graph.New(myApp)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(run(ctx, mux))

}

func run(ctx context.Context, mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:    net.JoinHostPort("0.0.0.0", fmt.Sprintf("%s", defaultPort)),
		Handler: mux,
	}

	errc := make(chan error, 1)
	go func() {
		errc <- srv.ListenAndServe()
	}()

	slog.Info("started", net.JoinHostPort("localhost", fmt.Sprintf("%s", defaultPort)))
	defer slog.Info("shutdown")

	var err error
	select {
	case err = <-errc:
	case <-ctx.Done():
		err = srv.Shutdown(context.Background())
	}

	if err != nil {
		return fmt.Errorf("srv.ListenAndServe: %w", err)
	}

	return nil
}

func configRead() config {
	file, err := os.Open("/build/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cfg := config{}
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
