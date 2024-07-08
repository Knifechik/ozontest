//go:build integration

package repo_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sipki-tech/database/connectors"
	"github.com/stretchr/testify/require"
	"log"
	"net"
	"ozon_test_compost/cmd/compost/internal/adapters/repo"
	"strconv"
	"testing"
	"time"
)

func start(t *testing.T) (context.Context, *repo.Repo, *require.Assertions) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=pass123",
			"POSTGRES_USER=user123",
			"POSTGRES_DB=postgres",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user123:pass123@%s/postgres?sslmode=disable", hostAndPort)

	log.Println("Connecting to database on url: ", databaseUrl)

	resource.Expire(60) // Tell docker to hard kill the container in 120 seconds

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		db, err := sql.Open("postgres", databaseUrl)
		if err != nil {
			t.Log(err)
			return err
		}
		err = db.Ping()
		if err != nil {
			t.Log(err)
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	//defer func() {
	//	if err := pool.Purge(resource); err != nil {
	//		log.Fatalf("Could not purge resource: %s", err)
	//	}
	//}()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	t.Cleanup(cancel)

	host, port, err := net.SplitHostPort(hostAndPort)
	if err != nil {
		t.Fatalf("Could not split host and port: %s", err)
	}
	portINT, err := strconv.Atoi(port)
	if err != nil {
		t.Fatalf("Could not split host and port: %s", err)
	}

	const migrateDir = "../../../migrate"
	connector := connectors.PostgresDB{
		User:     "user123",
		Password: "pass123",
		Host:     host,
		Port:     portINT,
		Database: "postgres",
		Parameters: &connectors.PostgresDBParameters{
			Mode: connectors.PostgresSSLDisable,
		},
	}
	dsn, err := connector.DSN()
	if err != nil {
		t.Fatalf("Could not dsn: %s", err)
	}

	cfg := repo.Config{
		Postgres:   repo.Connector{ConnectionDSN: dsn},
		MigrateDir: migrateDir,
		Driver:     "postgres",
	}

	repos, err := repo.New(ctx, cfg)
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	assert := require.New(t)

	time.Sleep(5 * time.Second)

	return context.Background(), repos, assert
}
